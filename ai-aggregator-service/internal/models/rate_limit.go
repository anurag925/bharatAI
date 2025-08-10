package models

import (
	"time"

	"github.com/google/uuid"
)

type RateLimit struct {
	ID             uuid.UUID `json:"id" db:"id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	ModelID        uuid.UUID `json:"model_id" db:"model_id"`
	Window         string    `json:"window" db:"window"`
	Limit          int       `json:"limit" db:"limit"`
	Used           int       `json:"used" db:"used"`
	ResetAt        time.Time `json:"reset_at" db:"reset_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type CreateRateLimit struct {
	OrganizationID uuid.UUID `json:"organization_id" validate:"required"`
	UserID         uuid.UUID `json:"user_id" validate:"required"`
	ModelID        uuid.UUID `json:"model_id" validate:"required"`
	Window         string    `json:"window" validate:"required,oneof=1m 1h 1d 30d"`
	Limit          int       `json:"limit" validate:"required,min=1"`
}

type UpdateRateLimit struct {
	Limit   *int       `json:"limit,omitempty" validate:"omitempty,min=1"`
	Used    *int       `json:"used,omitempty" validate:"omitempty,min=0"`
	ResetAt *time.Time `json:"reset_at,omitempty"`
}

type RateLimitResponse struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	ModelID        uuid.UUID `json:"model_id"`
	Window         string    `json:"window"`
	Limit          int       `json:"limit"`
	Used           int       `json:"used"`
	ResetAt        time.Time `json:"reset_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RateLimitWithDetails struct {
	RateLimitResponse
	Organization OrganizationResponse `json:"organization"`
	User         UserResponse         `json:"user"`
	Model        ModelResponse        `json:"model"`
}

type RateLimitCheck struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	ModelID        uuid.UUID `json:"model_id"`
	Window         string    `json:"window"`
	Allowed        bool      `json:"allowed"`
	Remaining      int       `json:"remaining"`
	ResetAt        time.Time `json:"reset_at"`
}
