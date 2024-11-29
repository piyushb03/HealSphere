package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/pawaspy/VitaReach/db/sqlc"
	"github.com/pawaspy/VitaReach/token"
	"github.com/pawaspy/VitaReach/util"
)

type createPatientRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Age      int32  `json:"age" binding:"required,gte=0"`
	Gender   string `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type patientResponse struct {
	Username  string             `json:"username"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Age       int32              `json:"age"`
	Gender    string             `json:"gender"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type loginPatientRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginPatientResponse struct {
	AccessToken string          `json:"access_token"`
	Patient     patientResponse `json:"patient"`
}

type updatePatientRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email" binding:"omitempty,email"`
	Phone  string `json:"phone"`
	Age    int32  `json:"age" binding:"omitempty,gte=0"`
	Gender string `json:"gender"`
}

type updatePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=6"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
}

func newPatientResponse(patient db.Patient) patientResponse {
	return patientResponse{
		Username:  patient.Username,
