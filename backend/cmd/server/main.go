package main

import (
	"log"
	"net/http"

	"github.com/atoscerebro/eviden-petshop/pkg/api"
	"github.com/atoscerebro/eviden-petshop/pkg/config"
	"github.com/atoscerebro/eviden-petshop/pkg/db"
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var err error
	// Load configuration
	config.LoadConfig()

	// Initialize the database connection
	dsn := config.AppConfig.DatabaseURI
	database, err := db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run the migrations
	db.Migrate(database.DB)

	// Initialize services
	petsService := &pets.Pets{DB: database}

	// Set up routes using chi router
	router := chi.NewRouter()

	// Use some basic middleware
	router.Use(middleware.Logger)

	// Set up API routes
	api.SetupRoutes(router, petsService)

	// Start the server on port 4000
	port := config.AppConfig.Port

	log.Printf("Starting server on port %s...", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

}
