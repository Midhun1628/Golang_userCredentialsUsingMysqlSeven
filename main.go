package main

import (
	"fmt"
	"log"
	"net/http"

	"Golang_userCredentialsUsingMysqlSeven/routes"
)

func main() {
	routes.InitializeRoutes()

	fmt.Println("Server is running at 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}


