package mocks

import (
	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"gorm.io/gorm"
)

// MockDB is a mock implementation of the Database interface
type MockDB struct {
	Pets []models.Pet
	Gorm *gorm.DB
}

// GetAllPets returns a list of mock pets
func (m *MockDB) GetAllPets() ([]models.Pet, error) {
	return m.Pets, nil
}

// CreatePet adds a new pet to the mock database
func (m *MockDB) CreatePet(pet *models.Pet) error {
	m.Pets = append(m.Pets, *pet)
	return nil
}

// Transaction is a mock implementation of the Transaction method from the db.Database interface.
func (m *MockDB) Transaction(fc func(tx *gorm.DB) error) error {
	// Create a new transaction.
	tx := m.Gorm.Begin()

	// Call the provided function with the transaction.
	err := fc(tx)

	// Roll back the transaction if there was an error.
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction.
	if err = tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
