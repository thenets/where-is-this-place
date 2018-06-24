package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


func main() {
	// Get arguments
	args := os.Args[1:]

	// Testing binary tree
	if args[0] == "test" {
		fmt.Println("Just testing...")
	}

	// Start web server
	if args[0] == "server" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path

			if path == "/" {
				json.NewEncoder(w).Encode("É comunista?")
			}
			if path == "/0" {
				json.NewEncoder(w).Encode("Resposta: Estados Unidos da América")
			}
			if path == "/1" {
				json.NewEncoder(w).Encode("Resposta: China")
			}
		})

		http.ListenAndServe(":5000", nil)
	}

}
