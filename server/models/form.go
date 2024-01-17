package models

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	FormID                  uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"form_id"`
	WorkspaceID             uuid.UUID               `gorm:"not null" json:"workspace_id"`
	Name                    string                  `gorm:"not null" json:"name,omitempty"`
	Path                    Path                    `gorm:"foreignKey:FormID" json:"path,omitempty"`
	Permits                 []Permit                `gorm:"foreignKey:FormID" json:"permits,omitempty"`
	DurationMeasurementUnit DurationMeasurementUnit `json:"duration_measurement_unit"`
	DurationLimit           uint8                   `json:"duration_limit"`
	DurationResetTime       string                  `json:"duration_reset_time"`
	CreatedAt               time.Time               `gorm:"not null" json:"created_at"`
	UpdatedAt               time.Time               `gorm:"not null" json:"updated_at"`
}

type DurationMeasurementUnit string

const (
	None    DurationMeasurementUnit = ""
	Minutes DurationMeasurementUnit = "minutes"
	Hours   DurationMeasurementUnit = "hours"
	Days    DurationMeasurementUnit = "days"
	Months  DurationMeasurementUnit = "months"
)

type FormCreateRequest struct {
	Name                    string                  `json:"name"`
	Path                    string                  `json:"path"`
	DurationMeasurementUnit DurationMeasurementUnit `json:"duration_measurement_unit"`
	DurationLimit           uint8                   `json:"duration_limit"`
	DurationResetTime       string                  `json:"duration_reset_time"`
}

type FormUpdateRequest struct {
	Name                    string                  `json:"name"`
	DurationMeasurementUnit DurationMeasurementUnit `json:"duration_measurement_unit"`
	DurationLimit           uint8                   `json:"duration_limit"`
	DurationResetTime       string                  `json:"duration_reset_time"`
}
