package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	userUuid := c.GetString("userUUID")
	user := database.User{}
	if err := h.DB.Where("uuid = ?", userUuid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	offers := []database.Offer{}
	h.DB.Where("user_id = ?", user.ID).Preload("EntryUsers").Find(&offers)

	offerResponses := []OfferGetResponse{}
	for _, offer := range offers {
		entryUsers := []UserGetResponse{}
		for _, user := range offer.EntryUsers {
			entryUsers = append(entryUsers, UserGetResponse{
				Uuid: user.UUID,
				Name: user.Name,
			})
		}
		offerResponses = append(offerResponses, OfferGetResponse{
			Uuid:        offer.UUID,
			Title:       offer.Title,
			Description: offer.Description,
			IsPublic:    offer.IsPublic,
			Deadline:    offer.Deadline,
			EntryUsers:  entryUsers,
		})
	}

	c.JSON(http.StatusOK, offerResponses)
}

func (h *OfferHandler) OffersUuidDelete(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	offerUuid := c.Param("uuid")

	if err := h.DB.Delete(
		&database.Offer{},
		"uuid = ? AND user_id = (SELECT id FROM users WHERE uuid = ?)",
		offerUuid,
		userUuid,
	).Error; err != nil {
		log.Printf("Error deleting offer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete offer", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer deleted successfully"})
}

func (h *OfferHandler) OffersUuidGet(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	offerUuid := c.Param("uuid")

	var offer database.Offer
	if err := h.DB.
		Preload("EntryUsers").
		Where("offers.uuid = ? AND users.uuid = ?", offerUuid, userUuid).
		Joins("join users on users.id = offers.user_id").
		First(&offer).Error; err != nil {
		log.Printf("Error retrieving offer: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Offer not found", "error": err.Error()})
		return
	}

	entryUsers := []UserGetResponse{}
	for _, user := range offer.EntryUsers {
		entryUsers = append(entryUsers, UserGetResponse{
			Uuid: user.UUID,
			Name: user.Name,
		})
	}
	offerResponse := OfferGetResponse{
		Uuid:        offer.UUID,
		Title:       offer.Title,
		Description: offer.Description,
		IsPublic:    offer.IsPublic,
		Deadline:    offer.Deadline,
		EntryUsers:  entryUsers,
	}

	c.JSON(http.StatusOK, offerResponse)
}

func (h *OfferHandler) OffersUuidPut(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	offerUuid := c.Param("uuid")

	var offerPostRequest OfferPostRequest
	if err := c.ShouldBindJSON(&offerPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var offer database.Offer
	if err := h.DB.
		Preload("EntryUsers").
		Where("offers.uuid = ? AND users.uuid = ?", offerUuid, userUuid).
		Joins("join users on users.id = offers.user_id").
		First(&offer).Error; err != nil {
		log.Printf("Error retrieving offer: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Offer not found", "error": err.Error()})
		return
	}

	updateData := database.Offer{
		Title:       offerPostRequest.Title,
		Description: offerPostRequest.Description,
		Giz:         uint(offerPostRequest.Giz),
		IsPublic:    offerPostRequest.IsPublic,
		Deadline:    offerPostRequest.Deadline,
	}
	if err := h.DB.Model(&offer).Updates(updateData).Error; err != nil {
		log.Printf("Error saving offer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update offer", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully"})
}

func (h *OfferHandler) OffersPost(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	user := database.User{}
	if err := h.DB.Where("uuid = ?", userUuid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var offerPostRequest OfferPostRequest
	if err := c.ShouldBindJSON(&offerPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// バリデーション
	giz := offerPostRequest.Giz
	// GIZのバリデーション
	// 1. GIZが0以上かチェック
	// 2. ユーザーがGIZを保有しているかチェック
	// 3. 保有していたらGIZを一時預かりユーザーに送金リクエスト
	// 4. 送金リクエストが成功したらオファー作成を続行

	chatURL := offerPostRequest.ChatUrl
	if chatURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Chat URL is required"})
		return
	}

	title := offerPostRequest.Title
	if title == "" || len(title) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Title is required and must be less than 255 characters"})
		return
	}

	description := offerPostRequest.Description
	if description == "" || len(description) > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Description is required and must be less than 1000 characters"})
		return
	}

	deadline := offerPostRequest.Deadline
	if deadline.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Deadline must be in the future"})
		return
	}

	offer := database.Offer{
		UUID:        uuid.New().String(),
		UserID:      user.ID,
		Giz:         uint(giz),
		ChatURL:     chatURL,
		Title:       title,
		Description: description,
		IsPublic:    offerPostRequest.IsPublic,
		Deadline:    deadline,
	}

	result := h.DB.Create(&offer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, OfferPostResponse{
		Uuid: offer.UUID,
	})
}

// // OffersUuidEntryDelete Delete /offers/:uuid/entry
// // Cancel an entry by UUID
//  OffersUuidEntryDelete(c *gin.Context)

// // OffersUuidEntryPost Post /offers/:uuid/entry
// // Entry an offer by UUID
//  OffersUuidEntryPost(c *gin.Context)

func (h *OfferHandler) OffersUuidEntryPost(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	offerUuid := c.Param("uuid")

	user := database.User{}
	if err := h.DB.Where("uuid = ?", userUuid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	offer := database.Offer{}
	if err := h.DB.Where("uuid = ?", offerUuid).First(&offer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Offer not found"})
		return
	}

	// TODO: doesn't work
	// offer.EntryUsersにuserが含まれているかチェック
	for _, entryUser := range offer.EntryUsers {
		if entryUser.ID == user.ID {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User already entered the offer"})
			return
		}
	}

	// エントリー作成
	entryUsers := append(offer.EntryUsers, &user)
	if err := h.DB.Model(&offer).Association("EntryUsers").Replace(entryUsers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create entry"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Entry created successfully"})
}

func (h *OfferHandler) OffersUuidEntryDelete(c *gin.Context) {
	userUuid := c.GetString("userUUID")
	offerUuid := c.Param("uuid")

	user := database.User{}
	if err := h.DB.Where("uuid = ?", userUuid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	offer := database.Offer{}
	if err := h.DB.Preload("EntryUsers").Where("uuid = ?", offerUuid).First(&offer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Offer not found"})
		return
	}

	if err := h.DB.Model(&offer).Association("EntryUsers").Delete(&user); err != nil {
		log.Printf("Error deleting entry: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete entry", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
