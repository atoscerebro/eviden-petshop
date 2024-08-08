package db

import (
	"log"
	"strings"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"gorm.io/gorm"
)

// GetAllPets retrieves all pets from the database, including their associated tags
func (g *GormDB) GetAllPets() ([]models.Pet, error) {
	var pets []models.Pet
	result := g.DB.Preload("Tags").Find(&pets)
	return pets, result.Error
}

// CreatePet adds a new pet to the database
func (g *GormDB) CreatePet(pet *models.Pet) error {
	// Start a transaction
	return g.DB.Transaction(func(tx *gorm.DB) error {
		// Create the pet record
		if err := tx.Create(pet).Error; err != nil {
			log.Printf("CreatePet: Failed to create pet: %v", err)
			return err
		}

		// Create the tag records and associate them with the pet
		for _, tag := range pet.Tags {
			if err := tx.Where(models.Tag{Name: tag.Name}).FirstOrCreate(&tag).Error; err != nil {
				return err
			}
			if err := tx.Model(pet).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}

		// Format the photo URLs as an array literal
		photoURLs := make([]string, len(pet.PhotoUrls))
		for i, url := range pet.PhotoUrls {
			photoURLs[i] = "'" + url + "'"
		}
		photoURLsStr := "{" + strings.Join(photoURLs, ",") + "}"

		// Update the pet record with the formatted photo URLs
		if err := tx.Model(pet).Update("photo_urls", photoURLsStr).Error; err != nil {
			return err
		}

		return nil
	})
}
