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
	VMake        string     `json:"v_make"`
	VModel       string     `json:"v_model"`
	VColor       string     `json:"v_color"`
	Expiry       *time.Time `json:"expiry"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updated_at"`
}

type PermitCreateRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PrimaryPhone string `json:"primary_phone"`
	VPlate       string `json:"v_plate"`
	VMake        string `json:"v_make"`
	VModel       string `json:"v_model"`
	VColor       string `json:"v_color"`
	Duration     int    `json:"duration"`
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
	Expiry       time.Time `gorm:"not null" json:"expiry"`
}
