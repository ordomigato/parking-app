package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ClientID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"client_id"`
	Username         string    `gorm:"uniqueIndex;not null" json:"username"`
	Password         string    `gorm:"not null" json:"password"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `gorm:"not null" json:"verified"`
	CreatedAt        time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time `gorm:"not null" json:"updated_at"`
	LastLogin        time.Time `gorm:"not null" json:"last_login"`
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

func FilterClientRecord(client *Client) ClientResponse {
	return ClientResponse{
		ClientID:  client.ClientID,
		Username:  client.Username,
		Verified:  client.Verified,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
}

func MigrateClient(db *gorm.DB) error {
	err := db.AutoMigrate(&Client{})
	return err
}
