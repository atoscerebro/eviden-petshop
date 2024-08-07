package api

import (
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux, petsService *pets.Pets) {
	router.Get("/pets", GetAllPetsHandler(petsService))
	router.Post("/pet", CreatePetHandler(petsService))

}
