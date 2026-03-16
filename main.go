package main

import (
	"fmt"
	"net/http"
	"os" // Required to read environment variables
)

func main() {
	// Get port from environment, or use 8080 as a default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8020" 
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! I am running on port %s", port)
	})

	fmt.Printf("Server starting on :%s...\n", port)
	// We add the colon here
	http.ListenAndServe(":"+port, nil)
}