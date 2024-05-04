package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/gizoffer/internal/database"
)

type UserHandler struct {
	DB *database.GizofferDB
}

func NewUserHandler(db *database.GizofferDB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) UsersGet(c *gin.Context) {
	// ここにGet all usersのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "All users fetched successfully"})
}

func (h *UserHandler) UsersIdDelete(c *gin.Context) {
	// ここにDelete a user by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) UsersIdGet(c *gin.Context) {
	// パスパラメータからIDを取得
	id := c.Param("id")

	// c.JSON(http.StatusOK, gin.H{"message": "User fetched successfully"})
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s fetched successfully", id)})
}

func (h *UserHandler) UsersIdPut(c *gin.Context) {
	// ここにUpdate a user by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) UsersPost(c *gin.Context) {
	// リクエストボディからユーザ情報を取得
	body := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
