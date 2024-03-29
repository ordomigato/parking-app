package models

import "github.com/google/uuid"

// TODO figure out foreign key syntax
type ClientWorkspace struct {
	WorkspaceID uuid.UUID `gorm:"type:uuid;primary_key" json:"workspace_id"`
	ClientID    uuid.UUID `gorm:"type:uuid;primary_key" json:"client_id"`
	RoleID      string    `json:"role_id,omitempty"`
}
