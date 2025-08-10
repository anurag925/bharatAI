package models

import (
	"time"

	"github.com/google/uuid"
)

type APIKey struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	UserID         uuid.UUID  `json:"user_id" db:"user_id"`
	OrganizationID uuid.UUID  `json:"organization_id" db:"organization_id"`
	KeyPrefix      string     `json:"key_prefix" db:"key_prefix"`
	KeyHash        string     `json:"-" db:"key_hash"`
	Name           string     `json:"name" db:"name"`
	Description    string     `json:"description" db:"description"`
	Permissions    JSONB      `json:"permissions" db:"permissions"`
	IsActive       bool       `json:"is_active" db:"is_active"`
	LastUsedAt     *time.Time `json:"last_used_at" db:"last_used_at"`
	ExpiresAt      *time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateAPIKey struct {
	Name        string     `json:"name" validate:"required,min=3,max=100"`
	Description string     `json:"description" validate:"max=500"`
	Permissions JSONB      `json:"permissions,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type UpdateAPIKey struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

type APIKeyResponse struct {
	ID          uuid.UUID  `json:"id"`
	KeyPrefix   string     `json:"key_prefix"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Permissions JSONB      `json:"permissions"`
	IsActive    bool       `json:"is_active"`
	LastUsedAt  *time.Time `json:"last_used_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type APIKeyWithSecret struct {
	APIKeyResponse
	Secret string `json:"secret"`
}
