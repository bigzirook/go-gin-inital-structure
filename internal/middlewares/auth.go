package middlewares

import (
	"net/http"
	"strings"

	"github.com/bigzirook/movie-ticket-booking/internal/services"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to protect routes and validate JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Validate the JWT token
		claims, err := services.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user claims in context for later use in controllers
		c.Set("user", claims)
		c.Next()
	}
}
