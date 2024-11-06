package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP server setup
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP!"))
	})

	log.Println("HTTP server listening on port 8000")
	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Set up the gRPC server (using a different port, e.g., 50051)
}
