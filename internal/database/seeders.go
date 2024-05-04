package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db *GizofferDB) UserSeed() {
	password1 := "password1"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPassword1 := string(bytes)

	password2 := "password2"
	bytes, err = bcrypt.GenerateFromPassword([]byte(password2), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPassword2 := string(bytes)

	users := []User{
		{
			UUID:           uuid.New().String(),
			Name:           "John Doe",
			Email:          "john@gmail.com",
			HashedPassword: hashedPassword1,
		},
		{
			UUID:           uuid.New().String(),
			Name:           "Jane Doe",
			Email:          "jane@gmail.com",
			HashedPassword: hashedPassword2,
		},
	}
	result := db.Create(&users)

	if result.Error != nil {
		log.Printf("failed to seed: %v", result.Error)
	} else {
		log.Printf("user seeded successfully")
	}
}

func (db *GizofferDB) OfferSeed() {
	offers := []Offer{
		{
			UUID:        uuid.New().String(),
			UserID:      1, // TODO: generate id in this method
			Giz:         100,
			ChatURL:     "https://chat.com",
			Title:       "Title",
			Description: "Description",
			IsPublic:    true,
			Deadline:    time.Now(),
		},
		{
			UUID:        uuid.New().String(),
			UserID:      2, // TODO: generate id in this method
			Giz:         200,
			ChatURL:     "https://chat.com",
			Title:       "Title",
			Description: "Description",
			IsPublic:    true,
			Deadline:    time.Now(),
		},
	}
	result := db.Create(&offers)

	if result.Error != nil {
		log.Printf("failed to seed: %v", result.Error)
	} else {
		log.Printf("offer seeded successfully")
	}
}

func (db *GizofferDB) EntrySeed() {
	entries := []Entry{
		{
			UUID:       uuid.New().String(),
			OfferID:    1, // TODO: generate id in this method
			UserID:     2, // TODO: generate id in this method
			IsApproved: true,
		},
		{
			UUID:       uuid.New().String(),
			OfferID:    2, // TODO: generate id in this method
			UserID:     1, // TODO: generate id in this method
			IsApproved: true,
		},
	}
	result := db.Create(&entries)

	if result.Error != nil {
		log.Printf("failed to seed: %v", result.Error)
	} else {
		log.Printf("entry seeded successfully")
	}
}

func (db *GizofferDB) Seed() {
	db.UserSeed()
	db.OfferSeed()
	db.EntrySeed()
}
