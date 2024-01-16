package models

import "github.com/google/uuid"

type Path struct {
	ID          uint      `gorm:"primary_key" json:"-"`
	Path        string    `gorm:"unique" json:"path"`
	WorkspaceID uuid.UUID `gorm:"not null" json:"-"`
	FormID      uuid.UUID `gorm:"not null" json:"-"`
}

type PathUpdateRequest struct {
	Path string `json:"path"`
}
