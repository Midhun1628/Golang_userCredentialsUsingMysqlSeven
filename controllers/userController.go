package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"Golang_userCredentialsUsingMysqlSeven/models" 

	"Golang_userCredentialsUsingMysqlSeven/utils.go"
)

// LoginHandler - Handles user login request
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.SendErrorResponse(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var input models.User
		contentType := r.Header.Get("Content-Type")

		if contentType == "application/json" {
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				utils.SendErrorResponse(w, "Invalid JSON format", http.StatusBadRequest)
				return
			}
		} else if contentType == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err != nil {
				utils.SendErrorResponse(w, "Invalid form data", http.StatusBadRequest)
				return
			}
			input.Username = r.FormValue("username")
			input.Password = r.FormValue("password")
		} else {
			utils.SendErrorResponse(w, "Unsupported content type", http.StatusUnsupportedMediaType)
			return
		}

		if input.Username == "" || input.Password == "" {
			utils.SendErrorResponse(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		isAuthenticated, err := models.Authenticate(db, input.Username, input.Password)
		if err != nil {
			log.Println("Error during authentication:", err)
			utils.SendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !isAuthenticated {
			utils.SendErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		utils.SendJSONResponse(w, map[string]string{"message": "User logged in successfully! Welcome " + input.Username})
	}
}

// GetUsersHandler - Retrieves all users
func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.SendErrorResponse(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		users, err := models.GetAllUsers(db)
		if err != nil {
			log.Println("Error fetching users:", err)
			utils.SendErrorResponse(w, "Failed to retrieve users", http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			utils.SendErrorResponse(w, "No users found", http.StatusNotFound)
			return
		}

		utils.SendJSONResponse(w, users)
	}
}
