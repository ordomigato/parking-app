package models

import (
	"time"

	"github.com/google/uuid"
)

type Blacklist struct {
	FormID    uuid.UUID `gorm:"type:uuid;not null;primary_key" json:"form_id"`
	VPlate    string    `gorm:"not null;primary_key" json:"v_plate"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type BlacklistCreateRequest struct {
	VPlate string `json:"v_plate"`
}
