package models

import (
	"time"

	"github.com/google/uuid"
)

type Permit struct {
	PermitId            uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"permit_id"`
	FirstName           string                   `gorm:"not null" json:"first_name"`
	LastName            string                   `gorm:"not null" json:"last_name"`
	Email               string                   `gorm:"not null" json:"email"`
	PrimaryPhone        string                   `gorm:"not null" json:"primary_phone"`
	VPlate              string                   `gorm:"not null" json:"v_plate"`
	VMake               string                   `gorm:"not null" json:"v_make"`
	VModel              string                   `gorm:"not null" json:"v_model"`
	VColor              string                   `gorm:"not null" json:"v_color"`
	Duration            uint8                    `gorm:"not null" json:"duration"`
	DurationMeasurement SubmissionConstraintType `gorm:"not null" json:"duration_measurement_type"`
	CreatedAt           time.Time                `gorm:"not null" json:"created_at"`
	UpdatedAt           time.Time                `gorm:"not null" json:"updated_at"`
}

type PermitCreateRequest struct {
	FirstName           string                   `json:"first_name"`
	LastName            string                   `json:"last_name"`
	Email               string                   `json:"email"`
	PrimaryPhone        string                   `json:"primary_phone"`
	VPlate              string                   `json:"v_plate"`
	VMake               string                   `json:"v_make"`
	VModel              string                   `json:"v_model"`
	VColor              string                   `json:"v_color"`
	Duration            uint8                    `json:"duration"`
	DurationMeasurement SubmissionConstraintType `json:"duration_measurement_type"`
}

type PermitUpdateRequest struct {
	FirstName           string                   `json:"first_name"`
	LastName            string                   `json:"last_name"`
	Email               string                   `json:"email"`
	PrimaryPhone        string                   `json:"primary_phone"`
	VPlate              string                   `json:"v_plate"`
	VMake               string                   `json:"v_make"`
	VModel              string                   `json:"v_model"`
	VColor              string                   `json:"v_color"`
	Duration            uint8                    `json:"duration"`
	DurationMeasurement SubmissionConstraintType `json:"duration_measurement_type"`
}
