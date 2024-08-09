package integration

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/atoscerebro/eviden-petshop/pkg/api"
	"github.com/atoscerebro/eviden-petshop/pkg/config"
	"github.com/atoscerebro/eviden-petshop/pkg/db"
	"github.com/atoscerebro/eviden-petshop/pkg/models"
	"github.com/atoscerebro/eviden-petshop/pkg/process/pets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lib/pq"
)

var router *chi.Mux
var database *db.Database

// TestMain is the entry point for testing. It sets up the environment before tests run.
func TestMain(m *testing.M) {
	var err error

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create database if it doesn't exist
	db.CreateDatabase(cfg)

	// Initialize the database connection
	var database *db.GormDB
	database, err = db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run the migrations
	db.Migrate(database.DB)

	// Initialize services
	petsService := &pets.Pets{DB: database}

	// Set up routes using chi router
	router = chi.NewRouter()

	// Use some basic middleware
	router.Use(middleware.Logger)

	// Set up API routes
	api.SetupRoutes(router, petsService)

	// Run tests
	code := m.Run()

	// Teardown logic (if necessary)
	// Example: database.DB.Close() if needed

	os.Exit(code)
}

func TestCreatePetIntegration(t *testing.T) {
	// Simulate an HTTP POST request to create a new pet
	pet := &models.Pet{
		Name:      "Buddy",
		PhotoUrls: pq.StringArray{"http://example.com/photo.jpg"},
	}
	jsonPet, _ := json.Marshal(pet)
	req, _ := http.NewRequest("POST", "/pet", bytes.NewBuffer(jsonPet))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req) // Use the router initialized in TestMain

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify the pet was actually created in the database

	// TODO
}
