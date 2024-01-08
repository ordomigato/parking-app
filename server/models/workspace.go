package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func MigrateWorkspace(db *gorm.DB) error {
	err := db.AutoMigrate(&Workspace{})
	return err
}
