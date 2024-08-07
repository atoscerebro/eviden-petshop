package db

import "github.com/atoscerebro/eviden-petshop/pkg/models"

// GetAllPets retrieves all pets from the database
func (g *GormDB) GetAllPets() ([]models.Pet, error) {
	var pets []models.Pet
	result := g.DB.Find(&pets)
	return pets, result.Error
}

// CreatePet adds a new pet to the database
func (g *GormDB) CreatePet(pet *models.Pet) error {
	result := g.DB.Create(pet)
	return result.Error
}
