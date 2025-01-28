package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/Okemwag/medihub/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

// AuthService provides methods for user authentication and token management.
type AuthService struct {
	db          *sql.DB          // Database connection
	jwtSecret   string           // Secret key for signing JWT tokens
	tokenExpiry time.Duration    // Expiry duration for JWT tokens
}

// NewAuthService creates a new instance of AuthService.
//
// @param jwtSecret string: The secret key used for signing JWT tokens.
// @param tokenExpiry time.Duration: The duration for which the JWT token is valid.
// @return *AuthService: A new AuthService instance.
func NewAuthService(jwtSecret string, tokenExpiry time.Duration) *AuthService {
	return &AuthService{
		db:          database.DB,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

// LoginResponse represents the response structure for a successful login.
type LoginResponse struct {
	Token string `json:"token"` // JWT token for authenticated user
}

// Login authenticates a user and generates a JWT token upon successful authentication.
//
// @param username string: The username of the user.
// @param password string: The password of the user.
// @return LoginResponse: The response containing the JWT token.
// @return error: An error if authentication fails or token generation fails.
func (s *AuthService) Login(username, password string) (LoginResponse, error) {
	var storedPassword string
	var role string
	var userID int64

	// Query the database for the user's credentials
	query := `SELECT id, password_hash, role_id FROM users WHERE username = $1`
	err := s.db.QueryRow(query, username).Scan(&userID, &storedPassword, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return LoginResponse{}, errors.New("invalid username or password")
		}
		return LoginResponse{}, errors.New("failed to authenticate: " + err.Error())
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return LoginResponse{}, errors.New("invalid username or password")
	}

	// Generate a JWT token for the authenticated user
	token, err := s.generateJWT(userID, role)
	if err != nil {
		return LoginResponse{}, errors.New("failed to generate token: " + err.Error())
	}

	// Return the token in the response
	return LoginResponse{Token: token}, nil
}

// Logout handles user logout. For stateless JWT, this typically involves client-side token invalidation.
//
// @param userID int64: The ID of the user logging out.
// @return error: An error if the logout process fails (not applicable for stateless JWT).
func (s *AuthService) Logout(userID int64) error {
	// For stateless JWT, logout would mean invalidating the token on the client side or maintaining a blacklist.
	// Implement logic to handle this if using session-based JWT (e.g., Redis blacklist).
	return nil
}

// generateJWT generates a JWT token for the given user ID and role.
//
// @param userID int64: The ID of the user.
// @param role string: The role of the user.
// @return string: The generated JWT token.
// @return error: An error if token generation fails.
func (s *AuthService) generateJWT(userID int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}