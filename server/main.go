package main

import (
	"log"
	"net/http"

	"Quortle/internal/api"
	"Quortle/internal/repository"
	"Quortle/internal/services"
)

func main() {
	repo := &repository.WordRepository{FilePath: "words.txt"}
	svc := services.NewWordService(repo)
	handler := api.NewHandler(svc)

	// Wrap handler with CORS
	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.Routes().ServeHTTP(w, r)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", corsHandler)
}
