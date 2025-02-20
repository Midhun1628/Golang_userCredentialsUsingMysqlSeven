package routes

import (
	"database/sql"
	"net/http"

	"Golang_userCredentialsUsingMysqlSeven/controllers"
)

// SetupRoutes initializes all API routes
func SetupRoutes(db *sql.DB) {
	http.HandleFunc("/login", controllers.LoginHandler(db))
	http.HandleFunc("/users", controllers.GetUsersHandler(db))
}
