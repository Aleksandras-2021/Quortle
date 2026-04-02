package api

import (
	"Quortle/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	service *services.WordService
}

func NewWordHandler(s *services.WordService) *WordHandler {
	return &WordHandler{service: s}
}

func (h *WordHandler) GetWordOfTheDay(c *gin.Context) {
	word, err := h.service.GetWordOfTheDay()
	if err != nil || word == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No word available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"word": word})
}

func (h *WordHandler) GetWordsTxt(c *gin.Context) {
	data, err := h.service.GetWordsTxt()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read words.txt"})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", data)
}
