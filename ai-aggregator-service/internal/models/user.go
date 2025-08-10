package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" db:"id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	Email          string    `json:"email" db:"email"`
	Username       string    `json:"username" db:"username"`
	FullName       string    `json:"full_name" db:"full_name"`
	AvatarURL      string    `json:"avatar_url" db:"avatar_url"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	IsSuperuser    bool      `json:"is_superuser" db:"is_superuser"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Metadata       JSONB     `json:"metadata" db:"metadata"`
}

type CreateUser struct {
	OrganizationID uuid.UUID `json:"organization_id" validate:"required"`
	Email          string    `json:"email" validate:"required,email,max=255"`
	Username       string    `json:"username" validate:"required,min=3,max=50,alphanum"`
	FullName       string    `json:"full_name" validate:"required,min=2,max=100"`
	AvatarURL      string    `json:"avatar_url" validate:"omitempty,url"`
	Metadata       JSONB     `json:"metadata,omitempty"`
}

type UpdateUser struct {
	Email     *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Username  *string `json:"username,omitempty" validate:"omitempty,min=3,max=50,alphanum"`
	FullName  *string `json:"full_name,omitempty" validate:"omitempty,min=2,max=100"`
	AvatarURL *string `json:"avatar_url,omitempty" validate:"omitempty,url"`
	IsActive  *bool   `json:"is_active,omitempty"`
	Metadata  JSONB   `json:"metadata,omitempty"`
}

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	AvatarURL      string    `json:"avatar_url"`
	IsActive       bool      `json:"is_active"`
	IsSuperuser    bool      `json:"is_superuser"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
