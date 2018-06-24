package binaryTree

import (
	"encoding/json"
	"net/http"
)

func main() {
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
