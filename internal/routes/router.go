package routes

import (
	"github.com/Okemwag/medihub/internal/controllers"
	"github.com/Okemwag/medihub/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the API routes for the application.
//
// This function defines the public and protected routes, including authentication and patient management endpoints.
// Protected routes require a valid JWT token, and some routes enforce role-based access control.
//
// @param router *gin.Engine: The Gin router instance.
// @param authController *controllers.AuthController: The controller for authentication-related endpoints.
// @param patientController *controllers.PatientController: The controller for patient-related endpoints.
// @param jwtSecret string: The secret key used for signing and validating JWT tokens.
func RegisterRoutes(router *gin.Engine, authController *controllers.AuthController, patientController *controllers.PatientController, jwtSecret string) {
	// Public Routes
	router.POST("/login", authController.Login)

	// Protected Routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Auth routes
		authGroup := protected.Group("/auth")
		{
			// Logout endpoint
			authGroup.POST("/logout", authController.Logout)
		}

		// Patient routes
		patientGroup := protected.Group("/patients")
		{
			// Create a new patient (only accessible to receptionists)
			patientGroup.POST("", middleware.RoleMiddleware("receptionist"), patientController.CreatePatient)

			// Update an existing patient (only accessible to receptionists)
			patientGroup.PUT("/:id", middleware.RoleMiddleware("receptionist"), patientController.UpdatePatient)

			// Delete a patient (only accessible to receptionists)
			patientGroup.DELETE("/:id", middleware.RoleMiddleware("receptionist"), patientController.DeletePatient)

			// Get a patient by ID (accessible to receptionists and doctors)
			patientGroup.GET("/:id", middleware.RoleMiddleware("receptionist", "doctor"), patientController.GetPatient)
		}
	}
}