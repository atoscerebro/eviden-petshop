package pets

import (
	"errors"

	"github.com/atoscerebro/eviden-petshop/pkg/db"
	"github.com/atoscerebro/eviden-petshop/pkg/models"
)

// Pets contains methods to interact with pets
type Pets struct {
	DB db.Database
}

// GetAllPets retrieves all pets using the provided database interface
func (pet *Pets) GetAllPets() ([]models.Pet, error) {
	return pet.DB.GetAllPets()
}

// CreatePet adds a new pet to the database
func (p *Pets) CreatePet(pet *models.Pet) error {
	if pet.Name == "" {
		return errors.New("pet name cannot be empty")
	}
	if len(pet.PhotoUrls) == 0 {
		return errors.New("photo URLs cannot be empty")
	}
	return p.DB.CreatePet(pet)
}
