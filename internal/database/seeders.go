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

	password3 := "password3"
	bytes, err = bcrypt.GenerateFromPassword([]byte(password3), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPassword3 := string(bytes)

	users := []User{
		{
			UUID:           uuid.New().String(),
			Name:           "John Doe",
			Email:          "john@gmail.com",
			HashedPassword: hashedPassword1,
			MyOffers: []*Offer{
				{
					UUID:        uuid.New().String(),
					UserID:      1,
					Giz:         100,
					ChatURL:     "https://chat.com",
					Title:       "Title",
					Description: "Description",
					IsPublic:    true,
					Deadline:    time.Now(),
					EntryUsers: []*User{
						{
							UUID:           uuid.New().String(),
							Name:           "Michael",
							Email:          "michael@gmail.com",
							HashedPassword: hashedPassword3,
						},
					},
				},
			},
			EntryOffers: []*Offer{
				{
					UUID:        uuid.New().String(),
					UserID:      2,
					Giz:         100,
					ChatURL:     "https://chat.com",
					Title:       "Title",
					Description: "Description",
					IsPublic:    true,
					Deadline:    time.Now(),
				},
			},
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
			Giz:         100,
			ChatURL:     "https://chat.com",
			Title:       "Title",
			Description: "Description",
			IsPublic:    true,
			Deadline:    time.Now(),
		},
		{
			UUID:        uuid.New().String(),
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

func (db *GizofferDB) Seed() {
	db.UserSeed()
}
