package main

import (
	"flag"
	"log"

	"github.com/iotassss/gizoffer/internal/database"
)

func main() {
	database.Connect()
	db := database.Connect()
	db.Refresh()
	log.Println("Database refresh completed.")

	seed := flag.Bool("seed", false, "Set to true to seed the database after migration")
	flag.Parse()

	if *seed {
		db.Seed()
		log.Println("Database seeding completed.")
	}
}
