package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GizofferDB struct {
	*gorm.DB
}

func Connect() *GizofferDB {
	// dsn := os.Getenv("MYSQL_USER") + ":" +
	// os.Getenv("DB_HOST") + ":" +
	// os.Getenv("DB_PORT") + ")/" +
	// dsn := "root:" +
	// 	os.Getenv("MYSQL_ROOT_PASSWORD") + "@tcp(" +
	// 	os.Getenv("DB_HOST") + ")/" +
	// 	os.Getenv("MYSQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:password@tcp(db)/gizoffer_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return &GizofferDB{db}
}
