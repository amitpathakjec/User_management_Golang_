package main

import (
	"log"
	"net/http"
	"user_management/db"
	"user_management/routes"
)

func main() {
	// Initializing the db connection
	db.InitDatabase()
	defer db.CloseDatabase()

	// user routes
	routes.SetupUserRoutes()

	// starting the server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
