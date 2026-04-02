package api

import (
	"Quortle/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	wordHandler *WordHandler
	userHandler *UserHandler
	authHandler *AuthHandler
}

func NewHandler(wordSvc *services.WordService, userSvc *services.UserService) *Handler {
	return &Handler{
		wordHandler: NewWordHandler(wordSvc),
		userHandler: NewUserHandler(userSvc),
		authHandler: NewAuthHandler(userSvc),
	}
}

// Routes returns a Gin Engine
func (h *Handler) Routes() *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", h.authHandler.Register)
		authRoutes.POST("/login", h.authHandler.Login)
	}

	r.GET("/word/random", h.wordHandler.GetWordOfTheDay)
	r.GET("/words.txt", h.wordHandler.GetWordsTxt)

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", h.userHandler.CreateUser)

		userRoutes.GET("/:username", h.userHandler.GetUser)
	}

	//r.GET("/word/secret", AuthMiddleware(), h.wordHandler.SecretWord) add middleware like this

	//Serve frontend files
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend" + c.Request.URL.Path)
	})
	return r
}
