package models

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	FormID                    uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"form_id"`
	WorkspaceID               uuid.UUID                `gorm:"not null" json:"workspace_id"`
	Name                      string                   `gorm:"not null" json:"name,omitempty"`
	SubmissionConstraintType  SubmissionConstraintType `gorm:"not null" json:"submission_constraint_type"`
	SubmissionConstraintLimit uint8                    `gorm:"not null" json:"submission_constraint_limit"`
	CreatedAt                 time.Time                `gorm:"not null" json:"created_at"`
	UpdatedAt                 time.Time                `gorm:"not null" json:"updated_at"`
}

type SubmissionConstraintType string

const (
	Minutes SubmissionConstraintType = "minutes"
	Hours   SubmissionConstraintType = "hours"
	Days    SubmissionConstraintType = "days"
	Months  SubmissionConstraintType = "months"
)

type FormCreateRequest struct {
	Name                      string                   `json:"name"`
	SubmissionConstraintType  SubmissionConstraintType `json:"submission_constraint_type"`
	SubmissionConstraintLimit uint8                    `json:"submission_constraint_limit"`
}
