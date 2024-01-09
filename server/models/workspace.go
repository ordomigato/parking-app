package models

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	WorkspaceID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"workspace_id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

type WorkspaceCreateRequest struct {
	Name string `json:"name"`
}

type WorkspaceResponse struct {
	WorkspaceID uuid.UUID `json:"workspace_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        string    `json:"role"`
}
