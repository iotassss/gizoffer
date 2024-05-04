package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/gizoffer/internal/database"
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
	c.JSON(http.StatusOK, gin.H{"message": "Auth successful"})
}
