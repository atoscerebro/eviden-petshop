package api

import (
	"encoding/json"
	"net/http"

	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
)

func GetAllPetsHandler(petsService *pets.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pets, err := petsService.GetAllPets()
		if err != nil {
			http.Error(w, "Failed to get pets", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(pets)
	}
}

// CreatePetHandler handles the request to create a new pet
func CreatePetHandler(petsService *pets.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pet models.Pet
		if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if err := petsService.CreatePet(&pet); err != nil {
			http.Error(w, "Failed to create pet", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pet)
	}
}
