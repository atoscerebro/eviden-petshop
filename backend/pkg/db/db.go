package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/atoscerebro/eviden-petshop/pkg/config"
	"github.com/atoscerebro/eviden-petshop/pkg/models"
	_ "github.com/lib/pq"
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
func InitDB(config *config.Config) (*GormDB, error) {

	dsn := fmt.Sprintf("%s/%s", config.DatabaseURI, config.DatabaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("InitDB: Failed to connect to database: %v", err)
		return nil, err
	}
	return &GormDB{DB: db}, nil
}

// creates a new database if it doesn't exist
func CreateDatabase(config *config.Config) {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres",
		config.DatabaseUser, config.DatabasePass, config.DatabaseHost, config.DatabasePort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("CreateDatabase: Failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.DatabaseName))
	if err != nil && err.Error() != fmt.Sprintf("pq: database \"%s\" already exists", config.DatabaseName) {
		log.Fatalf("Failed to create database: %v", err)
	}
}
