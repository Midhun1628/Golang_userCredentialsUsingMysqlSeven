// Separately  created files and folders In Golang 

package main

import (
	"fmt"
	"log"
	"net/http"

	"Golang_userCredentialsUsingMysqlSeven/config"
	"Golang_userCredentialsUsingMysqlSeven/routes"
)

func main() {
	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Setup routes
	routes.SetupRoutes(db)

	fmt.Println("Server is running at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
