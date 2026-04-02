package api

import (
	"Quortle/internal/auth"
	"Quortle/internal/services"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	user_service *services.UserService
}

func NewAuthHandler(u *services.UserService) *AuthHandler {
	return &AuthHandler{user_service: u}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.user_service.CreateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.user_service.GetUser(req.Username)
	if err != nil || !h.user_service.CheckPassword(user, req.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	secure := true

	if os.Getenv("LOCAL") == "true" {
		secure = false
	}
	c.SetCookie("token", token, 3600, "/", "", secure, true)

	c.JSON(200, gin.H{"message": "Logged in successfully"})
}
