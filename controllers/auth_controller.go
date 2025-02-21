package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"Golang_userCredentialsUsingMysqlSeven/config"
	"Golang_userCredentialsUsingMysqlSeven/models"
	"Golang_userCredentialsUsingMysqlSeven/utils.go"
)

// LoginHandler authenticates a user and returns a JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	                                         var input models.Credentials

	// Read data based on Content-Type
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		// Parse JSON request
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
	} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		// Parse form data request
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}
		input.Username = r.FormValue("username")
		input.Password = r.FormValue("password")
	} else {
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check credentials in the database
	var user models.Credentials
	err = db.QueryRow("SELECT username, password FROM credentials WHERE username=? AND password=?", input.Username, input.Password).
		Scan(&user.Username, &user.Password)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Construct response message
	message := "User logged in successfully! Welcome " + user.Username

	// Check Accept header to decide response format
	acceptHeader := r.Header.Get("Accept")

	if strings.Contains(acceptHeader, "application/x-www-form-urlencoded") {
		// URL-encoded response
		response := url.Values{}
		response.Set("message", message)
		response.Set("token", token)

		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Encode()))
	} else {
		// Default to JSON response
		response := map[string]string{
			"message": message,
			"token":   token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
