package app

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	log.Printf("Middleware for authentication is not implemented yet.")
	// 	c.Next()
	// }

	// WIP: Implement authentication middleware

	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, BearerSchema)

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
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

		c.Set("userID", claims.Subject) // ユーザーIDをコンテキストにセット
		c.Next()
	}
}
