package pets_test

import (
	"testing"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
	"github.com/atoscerebro/eviden-petshop/test/mocks"
)

func TestCreatePet(t *testing.T) {
	mockDB := &mocks.MockDB{
		Pets: []models.Pet{},
	}
	petsService := pets.Pets{DB: mockDB}

	newPet := &models.Pet{Name: "Doggie"}
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
