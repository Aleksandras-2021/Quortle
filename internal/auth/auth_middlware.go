package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authentication required"})
			return
		}

		claims, err := ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		newToken, _ := GenerateToken(claims.Username)
		c.SetCookie("token", newToken, 3600, "/", "", false, true)

		c.Set("username", claims.Username)
		c.Next()
	}
}
