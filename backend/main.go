package main

import (
	"backend/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Check if the connection is established by pinging the DB (optional)
	db, err := config.DB.DB()
	if err != nil {
		log.Fatal("Error getting database instance:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Continue with other setups like routes
	log.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
