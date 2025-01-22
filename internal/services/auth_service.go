package services

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/Okemwag/medihub/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db          *sql.DB
	jwtSecret   string
	tokenExpiry time.Duration
}

func NewAuthService(jwtSecret string, tokenExpiry time.Duration) *AuthService {
	return &AuthService{
		db:          database.DB,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	var storedPassword string
	var role string
	var userID int64

	query := `SELECT id, password_hash, role_id FROM users WHERE username = $1`
	err := s.db.QueryRow(query, username).Scan(&userID, &storedPassword, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("invalid username or password")
		}
		return "", errors.New("failed to authenticate: " + err.Error())
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := s.generateJWT(userID, role)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}

	return token, nil
}

func (s *AuthService) Logout(userID int64) error {
	// For stateless JWT, logout would mean invalidating the token on the client side or maintaining a blacklist.
	// Implement logic to handle this if using session-based JWT (e.g., Redis blacklist).
	return nil
}

func (s *AuthService) generateJWT(userID int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
