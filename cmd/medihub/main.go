package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/Okemwag/medihub/internal/controllers"
	"github.com/Okemwag/medihub/internal/routes"
	"github.com/Okemwag/medihub/internal/seeder"
	"github.com/Okemwag/medihub/internal/services"
	"github.com/Okemwag/medihub/pkg/config"
	"github.com/Okemwag/medihub/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
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

	// Run migrations using Goose
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

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow requests from this origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"}, // Exposed headers
		AllowCredentials: true, // Allow credentials (e.g., cookies)
	}))

	// Register routes
	routes.RegisterRoutes(router, authController, patientController, jwtSecret)

	// Start the server
	port := "8000"
	log.Printf("Server is starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// runMigrations runs database migrations using Goose.
func runMigrations(db *sql.DB, migrationsDir string) error {
	// Set the dialect for Goose (PostgreSQL in this case)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Run migrations
	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	log.Println("Migrations executed successfully.")
	return nil
}