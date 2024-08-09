package pets_test

import (
	"testing"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
	"github.com/atoscerebro/eviden-petshop/test/mocks"
	"github.com/lib/pq"
)

func TestCreatePet(t *testing.T) {
	mockDB := &mocks.MockDB{
		Pets: []models.Pet{},
	}
	petsService := pets.Pets{DB: mockDB}

	newPet := &models.Pet{
		Name:      "Doggie",
		PhotoUrls: pq.StringArray{"http://example.com/1.jpg", "http://example.com/2.jpg"},
	}
	err := petsService.CreatePet(newPet)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	result, err := petsService.GetAllPets()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 pet, got %d", len(result))
	}
	if result[0].Name != "Doggie" {
		t.Errorf("expected pet name NewPet, got %s", result[0].Name)
	}
}

func TestGetAllPets(t *testing.T) {
	mockDB := &mocks.MockDB{
		Pets: []models.Pet{
			{Name: "Bear"},
			{Name: "Milo"},
		},
	}
	petsService := pets.Pets{DB: mockDB}

	result, err := petsService.GetAllPets()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 pets, got %d", len(result))
	}
	if result[0].Name != "Bear" {
		t.Errorf("expected pet name MockPet1, got %s", result[0].Name)
	}
	if result[1].Name != "Milo" {
		t.Errorf("expected pet name MockPet2, got %s", result[1].Name)
	}
}

func TestCreatePet_EdgeCases(t *testing.T) {
	mockDB := &mocks.MockDB{
		Pets: []models.Pet{},
	}
	petsService := pets.Pets{DB: mockDB}

	// 1. Test creating a pet with an empty name
	petWithEmptyName := &models.Pet{
		Name:      "",
		PhotoUrls: pq.StringArray{"http://randompic.com/1.jpg"},
	}
	err := petsService.CreatePet(petWithEmptyName)
	if err == nil {
		t.Errorf("expected error when creating pet with empty name, got nil")
	}

	// 2. Test creating a pet with empty photo URL
	petWithInvalidURL := &models.Pet{
		Name:      "Fluffy",
		PhotoUrls: nil,
	}
	err = petsService.CreatePet(petWithInvalidURL)
	if err == nil {
		t.Errorf("expected error when creating pet with invalid photo URL, got nil")
	}
}
