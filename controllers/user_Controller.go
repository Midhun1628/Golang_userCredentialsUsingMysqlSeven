package controllers

import (
	
	"encoding/json"
	"log"
	"net/http"
	

	"Golang_userCredentialsUsingMysqlSeven/config"
	"Golang_userCredentialsUsingMysqlSeven/models"
	"Golang_userCredentialsUsingMysqlSeven/utils.go"
)

// GetUsersHandler retrieves all users (protected route)
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Extract the token (remove "Bearer ")
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}
	token := authHeader[len(bearerPrefix):] // Extract the actual token

	// Validate Token
	claims, valid := utils.ValidateToken(token)
	if !valid {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Extract username from claims
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}
	log.Println("Authenticated user:", username)

	// Database connection
	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Fetch users from MySQL
	rows, err := db.Query("SELECT username, password FROM credentials")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.Credentials
	for rows.Next() {
		var user models.Credentials
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			continue
		}
		users = append(users, user)
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
