package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Okemwag/medihub/internal/models"
)

type PatientService struct {
	db *sql.DB
}

func NewPatientService(db *sql.DB) *PatientService {
	return &PatientService{db: db}
}

func (s *PatientService) CreatePatient(ctx context.Context, patient *models.Patient) (int64, error) {
	query := `
		INSERT INTO patients (first_name, last_name, date_of_birth, gender, contact_number, email, address, medical_history, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`
	var id int64
	err := s.db.QueryRowContext(ctx, query,
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
	).Scan(&id)
	if err != nil {
		log.Printf("Error creating patient: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *PatientService) GetPatient(ctx context.Context, id int64) (*models.Patient, error) {
	query := `
		SELECT id, first_name, last_name, date_of_birth, gender, contact_number, email, address, medical_history, created_by, updated_by, created_at, updated_at
		FROM patients
		WHERE id = $1
	`
	var patient models.Patient
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.ContactNumber,
		&patient.Email,
		&patient.Address,
		&patient.MedicalHistory,
		&patient.CreatedBy,
		&patient.UpdatedBy,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("patient not found")
		}
		log.Printf("Error retrieving patient: %v", err)
		return nil, err
	}
	return &patient, nil
}

func (s *PatientService) UpdatePatient(ctx context.Context, id int64, patient *models.Patient) error {
	query := `
		UPDATE patients
		SET first_name = $1, last_name = $2, date_of_birth = $3, gender = $4, contact_number = $5, email = $6, address = $7, medical_history = $8, updated_by = $9, updated_at = CURRENT_TIMESTAMP
		WHERE id = $10
	`
	_, err := s.db.ExecContext(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.ContactNumber,
		patient.Email,
		patient.Address,
		patient.MedicalHistory,
		patient.UpdatedBy,
		id,
	)
	if err != nil {
		log.Printf("Error updating patient: %v", err)
		return err
	}
	return nil
}

func (s *PatientService) DeletePatient(ctx context.Context, id int64) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting patient: %v", err)
		return err
	}
	return nil
}
