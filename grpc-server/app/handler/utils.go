package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// verifyLogin is a middleware that checks if the user is logged in by verifying the auth token.
func (r *Router) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := getAuthToken(c)
		if t == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if _, err := r.grpcClient.VerifyToken(t); err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized", "message": err.Error()})
			c.Abort()
			return
		}
		c.Next() // Proceed to the next handler if the token is valid
	}
}

func getAuthToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	// Assuming the format is "Bearer <token>"
	split := strings.Split(authHeader, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		return ""
	}
	return split[1]
}
