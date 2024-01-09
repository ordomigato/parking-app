package models

import "github.com/google/uuid"

type Role struct {
	RoleID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"role_id"`
	RoleName string    `gorm:"uniqueIndex;not null" json:"name"`
}
