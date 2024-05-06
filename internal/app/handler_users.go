package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iotassss/gizoffer/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB *database.GizofferDB
}

func NewUserHandler(db *database.GizofferDB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) UsersUuidGet(c *gin.Context) {
	id := c.Param("id")

	// c.JSON(http.StatusOK, gin.H{"message": "User fetched successfully"})
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s fetched successfully", id)})
}

func (h *UserHandler) UsersPost(c *gin.Context) {
	var userPostRequest UserPostRequest
	if err := c.ShouldBindJSON(&userPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: validation

	// UUID
	uuid := uuid.New().String()

	// hashed password
	password := userPostRequest.Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	hashedPassword := string(bytes)

	user := database.User{
		UUID:           uuid,
		Email:          userPostRequest.Email,
		Name:           userPostRequest.Name,
		HashedPassword: hashedPassword,
	}

	result := h.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": uuid})
}

func (h *UserHandler) LoginPost(c *gin.Context) {
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

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
