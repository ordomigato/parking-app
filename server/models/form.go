package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Form struct {
	FormID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"form_id"`
	WorkspaceID uuid.UUID `gorm:"not null" json:"workspace_id"`
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Path        Path      `gorm:"foreignKey:FormID" json:"path,omitempty"`
	Permits     []Permit  `gorm:"foreignKey:FormID" json:"permits,omitempty"`
	// Note: example - 12 day, per 1 month
	CycleData CycleData `gorm:"type:jsonb" json:"cycle_data"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// Value Marshal
func (a CycleData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *CycleData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type CycleData struct {
	DurationLimit DurationLimit `json:"duration_limit"`
	ResetInterval ResetInterval `json:"reset_interval"`
}

type ResetInterval struct {
	Value   int                     `json:"value"`
	Unit    DurationMeasurementUnit `json:"unit"`
	RefDate time.Time               `json:"reference_date"`
}

type DurationLimit struct {
	Unit  DurationMeasurementUnit `json:"unit"`
	Value int                     `json:"value"`
}

type DurationMeasurementUnit string

const (
	None    DurationMeasurementUnit = ""
	Minutes DurationMeasurementUnit = "minutes"
	Hours   DurationMeasurementUnit = "hours"
	Days    DurationMeasurementUnit = "days"
	Months  DurationMeasurementUnit = "months"
	Years   DurationMeasurementUnit = "years"
)

type FormCreateRequest struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CycleData CycleData `json:"cycle_data"`
}

type FormUpdateRequest struct {
	Name      string    `json:"name"`
	CycleData CycleData `json:"cycle_data"`
}
