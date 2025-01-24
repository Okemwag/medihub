package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Okemwag/medihub/internal/controllers"
	"github.com/Okemwag/medihub/internal/routes"
	"github.com/Okemwag/medihub/internal/seeder"
	"github.com/Okemwag/medihub/internal/services"
	"github.com/Okemwag/medihub/pkg/config"
	"github.com/Okemwag/medihub/pkg/database"
	"github.com/gin-gonic/gin"
)

// @title Medihub API
// @version 1.0
// @description This is the API documentation for the Medihub application.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@medihub.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database.InitDB(cfg)
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	// Run migrations
	if err := runMigrations(database.DB, "./migrations"); err != nil {
		log.Printf("Warning: Error running migrations: %v", err)
		log.Println("Proceeding to start the server...")
	}

	// Seed database
	log.Println("Seeding database...")
	seeder.SeedUsers()

	// Initialize AuthService and AuthController
	jwtSecret := os.Getenv("JWT_SECRET") // Retrieve JWT secret from environment variables
	tokenExpiry := 24 * time.Hour
	authService := services.NewAuthService(jwtSecret, tokenExpiry)
	authController := controllers.NewAuthController(authService)

	// Initialize PatientController
	patientController := controllers.NewPatientController(database.DB)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, authController, patientController, jwtSecret)

	// Start the server
	port := "8000"
	log.Printf("Server is starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// runMigrations reads and executes .sql files from the migrations folder.
func runMigrations(db *sql.DB, migrationsDir string) error {
	// Read all files in the migrations directory
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	// Filter and sort .sql files
	var sqlFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, filepath.Join(migrationsDir, file.Name()))
		}
	}
	sort.Strings(sqlFiles)

	// Execute each SQL file
	for _, file := range sqlFiles {
		log.Printf("Executing migration: %s", file)
		sqlContent, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		// Execute SQL content, log non-critical errors
		if _, err := db.Exec(string(sqlContent)); err != nil {
			if strings.Contains(err.Error(), "relation already exists") {
				log.Printf("Skipping migration %s: %v", file, err)
				continue
			}
			return err
		}
	}

	log.Println("Migrations executed successfully.")
	return nil
}