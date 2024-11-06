package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HTTP server setup
	http.HandleFunc("/", handler)
	log.Println("HTTP server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Budget Tracker (^-^)")
}
