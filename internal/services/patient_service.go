package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/Okemwag/medihub/internal/models"

)

type PatientService struct {
	DB *sql.DB
}

func NewPatientService(db *sql.DB) *PatientService {
	return &PatientService{DB: db}
}

func (s *PatientService) CreatePatient(ctx context.Context, patient *models.Patient) error {
	query := `
		INSERT INTO patients (first_name, last_name, date_of_birth, gender, contact_number, email, address, medical_history, created_by, updated_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`
	if patient.CreatedAt.IsZero() {
		patient.CreatedAt = time.Now()
	}

	if patient.UpdatedAt.IsZero() {
		patient.UpdatedAt = time.Now()
	}

	err := s.DB.QueryRowContext(
		ctx,
		query,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.ContactNumber,
		patient.Email,
		patient.Address,
		patient.MedicalHistory,
		patient.CreatedBy,
		patient.UpdatedBy,
		patient.CreatedAt,
		patient.UpdatedAt,
	).Scan(&patient.ID)

	if err != nil {
		return fmt.Errorf("error creating patient: %v", err)
	}

	log.Printf("Created patient with ID: %d", patient.ID)
	return nil
}
