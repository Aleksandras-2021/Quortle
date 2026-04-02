package api

import (
	"Quortle/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	wordHandler *WordHandler
	userHandler *UserHandler
}

func NewHandler(wordSvc *services.WordService, userSvc *services.UserService) *Handler {
	return &Handler{
		wordHandler: NewWordHandler(wordSvc),
		userHandler: NewUserHandler(userSvc),
	}
}

// Routes returns a Gin Engine
func (h *Handler) Routes() *gin.Engine {
	r := gin.Default()

	// Serve frontend files
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	// Word routes
	r.GET("/word/random", h.wordHandler.GetWordOfTheDay)
	r.GET("/words.txt", h.wordHandler.GetWordsTxt)

	// User routes
	userRoutes := r.Group("/users")
	{
		// POST /users
		userRoutes.POST("", h.userHandler.CreateUser)

		// GET /users/:username
		userRoutes.GET("/:username", h.userHandler.GetUser)
	}

	return r
}
