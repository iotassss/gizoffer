package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/gizoffer/internal/database"
)

type EntryHandler struct {
	DB *database.GizofferDB
}

func NewEntryHandler(db *database.GizofferDB) *EntryHandler {
	return &EntryHandler{
		DB: db,
	}
}

func (h *EntryHandler) EntriesIdDelete(c *gin.Context) {
	// ここにDelete an entry by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}

func (h *EntryHandler) EntriesIdGet(c *gin.Context) {
	// ここにGet an entry by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "Entry fetched successfully"})
}

func (h *EntryHandler) EntriesPost(c *gin.Context) {
	// ここにCreate an entryのビジネスロジックを実装
	c.JSON(http.StatusCreated, gin.H{"message": "Entry created successfully"})
}
