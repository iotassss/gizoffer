package app

import (
	"fmt"
	"net/http"

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

	c.JSON(http.StatusCreated, gin.H{"user_id": uuid})
}
