package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	DisplayName string    `json:"display_name" db:"display_name"`
	Description string    `json:"description" db:"description"`
	Website     string    `json:"website" db:"website"`
	LogoURL     string    `json:"logo_url" db:"logo_url"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Metadata    JSONB     `json:"metadata" db:"metadata"`
}

type CreateOrganization struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Website     string `json:"website" validate:"omitempty,url"`
	LogoURL     string `json:"logo_url" validate:"omitempty,url"`
	Metadata    JSONB  `json:"metadata,omitempty"`
}

type UpdateOrganization struct {
	DisplayName *string `json:"display_name,omitempty" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
	Website     *string `json:"website,omitempty" validate:"omitempty,url"`
	LogoURL     *string `json:"logo_url,omitempty" validate:"omitempty,url"`
	IsActive    *bool   `json:"is_active,omitempty"`
	Metadata    JSONB   `json:"metadata,omitempty"`
}

type OrganizationResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	LogoURL     string    `json:"logo_url"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
