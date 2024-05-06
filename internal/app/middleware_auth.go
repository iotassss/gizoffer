package app

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// WIP: Implement authentication middleware

	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, BearerSchema)

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// TODO:
		// Check if the user exists in the database
		// Check if the user is active
		// Check if Issuer is correct
		// Check if ExpiresAt is in the future

		c.Set("userUUID", claims.Subject) // ユーザーIDをコンテキストにセット
		c.Next()
	}
}
