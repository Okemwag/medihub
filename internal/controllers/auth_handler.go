package controllers

import (
	"net/http"

	"github.com/Okemwag/medihub/internal/services"
	"github.com/gin-gonic/gin"
)

// AuthController handles HTTP requests related to user authentication.
type AuthController struct {
	authService *services.AuthService // Service for authentication-related operations
}

// NewAuthController creates a new instance of AuthController.
//
// @param authService *services.AuthService: The authentication service.
// @return *AuthController: A new AuthController instance.
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login authenticates a user and returns a JWT token upon successful authentication.
//
// @Summary Authenticate a user
// @Description Authenticate a user with the provided username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body struct{Username string; Password string} true "Login credentials"
// @Success 200 {object} services.LoginResponse "Returns the JWT token"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 401 {object} map[string]string "Invalid username or password"
// @Router /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind the request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Authenticate the user and generate a JWT token
	response, err := ctrl.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, response)
}

// Logout handles user logout. For stateless JWT, this typically involves client-side token invalidation.
//
// @Summary Logout a user
// @Description Logout the currently authenticated user
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string "Confirmation message"
// @Router /logout [post]
func (ctrl *AuthController) Logout(c *gin.Context) {
	// Add session handling if needed
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}