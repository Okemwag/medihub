package models

import "time"

type Patient struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	ContactNumber  string    `json:"contact_number"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	MedicalHistory string    `json:"medical_history"`
	CreatedBy      int64     `json:"created_by"`
	UpdatedBy      int64     `json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

