package api

import (
	"Quortle/internal/services"
	"encoding/json"
	"net/http"
	"os"
)

type Handler struct {
	service *services.WordService
}

func NewHandler(s *services.WordService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/word/random", h.GetWordOfTheDay)

	mux.HandleFunc("/words.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		data, err := os.ReadFile("words.txt")
		if err != nil {
			http.Error(w, "Unable to read words.txt", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	mux.Handle("/", http.FileServer(http.Dir("./frontend")))

	return mux
}

func (h *Handler) GetWordOfTheDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	word, err := h.service.GetWordOfTheDay()
	if err != nil || word == "" {
		http.Error(w, "No word available", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"word": word,
	})
}
