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

	mux := handler.Routes() // create once

	// Wrap mux with CORS
	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r) // reuse mux
	})

	log.Println("Server running on :80")
	http.ListenAndServe(":80", corsHandler)
}
