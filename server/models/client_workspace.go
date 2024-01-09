package models

import "github.com/google/uuid"

type ClientWorkspace struct {
	WorkspaceID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"workspace_id"`
	ClientID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"client_id"`
	RoleID      string    `gorm:"not null" json:"role_id,omitempty"`
}
