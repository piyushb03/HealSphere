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

type createDoctorRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	Qualification  string `json:"qualification" binding:"required"`
	Experience     int32  `json:"experience" binding:"required,gte=0"`
	Password       string `json:"password" binding:"required,min=6"`
}

type doctorResponse struct {
	Username       string             `json:"username"`
	Name           string             `json:"name"`
	Email          string             `json:"email"`
	Phone          string             `json:"phone"`
	Gender         string             `json:"gender"`
	Specialization string             `json:"specialization"`
	Qualification  string             `json:"qualification"`
	Experience     int32              `json:"experience"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}

type loginDoctorRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginDoctorResponse struct {
	AccessToken string         `json:"access_token"`
	Doctor      doctorResponse `json:"doctor"`
}

type updateDoctorRequest struct {
	Name           string `json:"name"`
