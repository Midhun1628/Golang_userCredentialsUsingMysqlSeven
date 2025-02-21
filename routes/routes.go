package routes

import (
	"net/http"

	"Golang_userCredentialsUsingMysqlSeven/controllers"
)

// InitializeRoutes sets up the HTTP routes
func InitializeRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/users", controllers.GetUsersHandler)
}

// SetupRoutes initializes all API routes
