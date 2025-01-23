package routes

import (
	"github.com/Okemwag/medihub/internal/controllers"
	"github.com/Okemwag/medihub/internal/middleware"
	"github.com/gin-gonic/gin"
)

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
			authGroup.POST("/logout", authController.Logout)
		}

		// Patient routes
		patientGroup := protected.Group("/patients")
		{
			// Only receptionists can create, update, or delete patients
			patientGroup.POST("", middleware.RoleMiddleware("receptionist"), patientController.CreatePatient)
			patientGroup.PUT("/:id", middleware.RoleMiddleware("receptionist"), patientController.UpdatePatient)
			patientGroup.DELETE("/:id", middleware.RoleMiddleware("receptionist"), patientController.DeletePatient)

			// Both receptionists and doctors can view patients
			patientGroup.GET("/:id", middleware.RoleMiddleware("receptionist", "doctor"), patientController.GetPatient)
		}
	}
}
