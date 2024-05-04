package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/gizoffer/internal/database"
)

type OfferHandler struct {
	DB *database.GizofferDB
}

func NewOfferHandler(db *database.GizofferDB) *OfferHandler {
	return &OfferHandler{
		DB: db,
	}
}

func (h *OfferHandler) OffersGet(c *gin.Context) {
	// ここにGet all offersのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "All offers fetched successfully"})
}

func (h *OfferHandler) OffersIdDelete(c *gin.Context) {
	// ここにDelete an offer by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "Offer deleted successfully"})
}

func (h *OfferHandler) OffersIdGet(c *gin.Context) {
	// ここにGet an offer by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "Offer fetched successfully"})
}

func (h *OfferHandler) OffersIdPut(c *gin.Context) {
	// ここにUpdate an offer by IDのビジネスロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully"})
}

func (h *OfferHandler) OffersPost(c *gin.Context) {
	// ここにCreate an offerのビジネスロジックを実装
	c.JSON(http.StatusCreated, gin.H{"message": "Offer created successfully"})
}
