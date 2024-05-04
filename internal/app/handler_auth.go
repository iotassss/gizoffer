package app

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iotassss/gizoffer/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *database.GizofferDB
}

func NewAuthHandler(db *database.GizofferDB) *AuthHandler {
	return &AuthHandler{
		DB: db,
	}
}

func (h *AuthHandler) LoginPost(c *gin.Context) {
	var loginPostRequest LoginPostRequest
	if err := c.ShouldBindJSON(&loginPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	var user database.User
	result := h.DB.Model(&database.User{}).Where("Email = ?", loginPostRequest.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginPostRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	token, err := generateToken(user.UUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// TODO: Check if this implementation is correct
func generateToken(userID string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // トークンの有効期限
		Issuer:    "gizoffer",                            // 発行者
		Subject:   userID,                                // ユーザーUUID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("your_secret_key"))
}
