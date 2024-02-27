package models

import (
	"time"

	"github.com/google/uuid"
)

type Permit struct {
	PermitID     uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"permit_id"`
	FormID       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"form_id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	PrimaryPhone string     `json:"primary_phone"`
	VPlate       string     `gorm:"not null" json:"v_plate"`
	VMake        string     `gorm:"column:v_make" json:"v_make"`
	VModel       string     `gorm:"column:v_model" json:"v_model"`
	VColor       string     `gorm:"column:v_color" json:"v_color"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updated_at"`
}

type PermitCreateRequest struct {
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	PrimaryPhone string     `json:"primary_phone"`
	VPlate       string     `json:"v_plate"`
	VMake        string     `json:"v_make"`
	VModel       string     `json:"v_model"`
	VColor       string     `json:"v_color"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
}

type PermitUpdateRequest struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PrimaryPhone string    `json:"primary_phone"`
	VPlate       string    `json:"v_plate"`
	VMake        string    `json:"v_make"`
	VModel       string    `json:"v_model"`
	VColor       string    `json:"v_color"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
}
