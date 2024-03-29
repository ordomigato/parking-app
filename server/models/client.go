package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ClientID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"client_id"`
	Username         string    `gorm:"uniqueIndex;not null" json:"username"`
	Password         string    `gorm:"not null" json:"password,omitempty"`
	VerificationCode string    `gorm:"null" json:"verification_code"`
	Verified         bool      `gorm:"not null" json:"verified"`
	CreatedAt        time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time `gorm:"not null" json:"updated_at"`
	LastLogin        time.Time `gorm:"not null" json:"last_login"`
}

type ClientResponse struct {
	ClientID  uuid.UUID `json:"client_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Verified  bool      `json:"verified,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
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

type VerifyClientRequest struct {
	Username string `json:"username"`
	OTP      string `json:"otp"`
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
