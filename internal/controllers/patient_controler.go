package controllers

import (
	"net/http"
	"strconv"

	"database/sql"
	"github.com/Okemwag/medihub/internal/models"
	"github.com/Okemwag/medihub/internal/services"
	"github.com/gin-gonic/gin"
)

// PatientController handles HTTP requests related to patient management.
type PatientController struct {
	patientService *services.PatientService // Service for patient-related operations
}

// NewPatientController creates a new instance of PatientController.
//
// @param db *sql.DB: A database connection.
// @return *PatientController: A new PatientController instance.
func NewPatientController(db *sql.DB) *PatientController {
	return &PatientController{
		patientService: services.NewPatientService(db),
	}
}

// CreatePatient handles the creation of a new patient.
//
// @Summary Create a new patient
// @Description Create a new patient with the input payload
// @Tags patients
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} map[string]int "Returns the ID of the created patient"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 401 {object} map[string]string "Unauthorized: User ID not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /patients [post]
func (c *PatientController) CreatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve user ID from the context (set during authentication)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found"})
		return
	}
	patient.CreatedBy = userID.(int64)
	patient.UpdatedBy = userID.(int64)

	// Create the patient using the service
	id, err := c.patientService.CreatePatient(ctx.Request.Context(), &patient)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetPatient retrieves a patient by ID.
//
// @Summary Get a patient by ID
// @Description Retrieve a patient record by its ID
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient "The patient record"
// @Failure 400 {object} map[string]string "Invalid patient ID"
// @Failure 404 {object} map[string]string "Patient not found"
// @Router /patients/{id} [get]
func (c *PatientController) GetPatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	// Retrieve the patient using the service
	patient, err := c.patientService.GetPatient(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

// UpdatePatient updates an existing patient by ID.
//
// @Summary Update a patient by ID
// @Description Update an existing patient record by its ID
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Updated patient data"
// @Success 204 "No content"
// @Failure 400 {object} map[string]string "Invalid patient ID or request payload"
// @Failure 401 {object} map[string]string "Unauthorized: User ID not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /patients/{id} [put]
func (c *PatientController) UpdatePatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve user ID from the context (set during authentication)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found"})
		return
	}
	patient.UpdatedBy = userID.(int64)

	// Update the patient using the service
	if err := c.patientService.UpdatePatient(ctx.Request.Context(), id, &patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeletePatient deletes a patient by ID.
//
// @Summary Delete a patient by ID
// @Description Delete a patient record by its ID
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 204 "No content"
// @Failure 400 {object} map[string]string "Invalid patient ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /patients/{id} [delete]
func (c *PatientController) DeletePatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	// Delete the patient using the service
	if err := c.patientService.DeletePatient(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}