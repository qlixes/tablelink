package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres dbname=tablelink password=postgres port=5432 sslmode=disable"), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
