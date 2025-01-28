package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/Okemwag/medihub/pkg/database"
)

func SeedUsers() {
	db := database.DB

	users := []struct {
		Username string
		Password string
		RoleID   int64
	}{
		{"admin", "@Doktari123", 1},
		{"receptionist", "@#PaSSwords123", 2},
	}

	for _, user := range users {
		// Check if the user already exists
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)
		if err != nil {
			log.Fatalf("failed to check if user exists: %v", err)
		}

		if exists {
			log.Printf("user %s already exists, skipping insertion", user.Username)
			continue
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}

		// Insert the user
		query := `INSERT INTO users (username, password_hash, role_id, created_at, updated_at)
			      VALUES ($1, $2, $3, $4, $5)`

		_, err = db.Exec(query, user.Username, string(hashedPassword), user.RoleID, time.Now(), time.Now())
		if err != nil {
			log.Fatalf("failed to insert user: %v", err)
		}

		log.Printf("user %s seeded successfully", user.Username)
	}
}