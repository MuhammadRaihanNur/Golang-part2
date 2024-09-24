package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler untuk merespons permintaan HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a simple Go HTTP server!")
}

func main() {
	http.HandleFunc("/", handler)

	// Jalankan server pada port 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
