package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ClientID         string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username         string `gorm:"uniqueIndex;not null"`
	Password         string `gorm:"not null"`
	VerificationCode string
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	LastLogin        time.Time `gorm:"not null"`
}

type ClientResponse struct {
	ClientID  uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"name,omitempty"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientRegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ClientLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func MigrateClient(db *gorm.DB) error {
	err := db.AutoMigrate(&Client{})
	return err
}
