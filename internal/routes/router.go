package routes

import (
	"github.com/Okemwag/medihub/internal/controllers"
	"github.com/Okemwag/medihub/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, authController *controllers.AuthController, jwtSecret string) {
	// Public Routes
	router.POST("/login", authController.Login)

	// Protected Routes
	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		protected.POST("/logout", authController.Logout)
	}
}
