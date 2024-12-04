package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/pawaspy/VitaReach/db/sqlc"
	"github.com/pawaspy/VitaReach/token"
)

// createAppointmentRequest defines the request parameters for creating an appointment
type createAppointmentRequest struct {
	DoctorUsername  string `json:"doctor_username" binding:"required"`
	DoctorName      string `json:"doctor_name" binding:"required"`
	AppointmentDate string `json:"appointment_date" binding:"required"`
	AppointmentTime string `json:"appointment_time" binding:"required"`
	Specialty       string `json:"specialty" binding:"required"`
	Symptoms        string `json:"symptoms" binding:"required"`
}

// appointmentResponse defines the response structure for appointment data
type appointmentResponse struct {
	ID              int64     `json:"id"`
	PatientUsername string    `json:"patient_username"`
	PatientName     string    `json:"patient_name,omitempty"`
	DoctorUsername  string    `json:"doctor_username"`
	DoctorName      string    `json:"doctor_name"`
	AppointmentDate string    `json:"appointment_date"`
	AppointmentTime string    `json:"appointment_time"`
	Specialty       string    `json:"specialty"`
	Symptoms        string    `json:"symptoms"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes,omitempty"`
	IsOnline        bool      `json:"is_online"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// newAppointmentResponse converts a db.Appointment to an appointmentResponse
func newAppointmentResponse(appointment db.Appointment) appointmentResponse {
	// Use YYYY-MM-DD format for date
	dateStr := ""
	if appointment.AppointmentDate.Valid {
		// Format time value extracted from pgtype.Date as YYYY-MM-DD
		dateStr = time.Date(
