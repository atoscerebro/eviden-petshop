package db

import (
	"log"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"gorm.io/gorm"
)

// Migrate runs the database migrations
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Pet{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
