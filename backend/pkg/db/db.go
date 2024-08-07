package db

import (
	"log"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database is an interface for database operations
type Database interface {
	GetAllPets() ([]models.Pet, error)
	CreatePet(pet *models.Pet) error
}

// GormDB is the real database implementation using GORM
type GormDB struct {
	DB *gorm.DB
}

// InitDB initializes the database connection
func InitDB(dsn string) (*GormDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	return &GormDB{DB: db}, nil
}
