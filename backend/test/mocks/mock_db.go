package mocks

import "github.com/atoscerebro/eviden-petshop/pkg/models"

// MockDB is a mock implementation of the Database interface
type MockDB struct {
	Pets []models.Pet
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
