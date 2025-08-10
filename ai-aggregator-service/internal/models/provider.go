package models

import (
	"time"

	"github.com/google/uuid"
)

type Provider struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	DisplayName      string    `json:"display_name" db:"display_name"`
	Description      string    `json:"description" db:"description"`
	LogoURL          string    `json:"logo_url" db:"logo_url"`
	Website          string    `json:"website" db:"website"`
	DocumentationURL string    `json:"documentation_url" db:"documentation_url"`
	BaseURL          string    `json:"base_url" db:"base_url"`
	APIVersion       string    `json:"api_version" db:"api_version"`
	IsActive         bool      `json:"is_active" db:"is_active"`
	Config           JSONB     `json:"config" db:"config"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProvider struct {
	Name             string `json:"name" validate:"required,min=2,max=50"`
	DisplayName      string `json:"display_name" validate:"required,min=2,max=100"`
	Description      string `json:"description" validate:"max=500"`
	LogoURL          string `json:"logo_url" validate:"omitempty,url"`
	Website          string `json:"website" validate:"omitempty,url"`
	DocumentationURL string `json:"documentation_url" validate:"omitempty,url"`
	BaseURL          string `json:"base_url" validate:"required,url"`
	APIVersion       string `json:"api_version" validate:"required,max=20"`
	Config           JSONB  `json:"config,omitempty"`
}

type UpdateProvider struct {
	DisplayName      *string `json:"display_name,omitempty" validate:"omitempty,min=2,max=100"`
	Description      *string `json:"description,omitempty" validate:"omitempty,max=500"`
	LogoURL          *string `json:"logo_url,omitempty" validate:"omitempty,url"`
	Website          *string `json:"website,omitempty" validate:"omitempty,url"`
	DocumentationURL *string `json:"documentation_url,omitempty" validate:"omitempty,url"`
	BaseURL          *string `json:"base_url,omitempty" validate:"omitempty,url"`
	APIVersion       *string `json:"api_version,omitempty" validate:"omitempty,max=20"`
	IsActive         *bool   `json:"is_active,omitempty"`
	Config           JSONB   `json:"config,omitempty"`
}

type ProviderResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	DisplayName      string    `json:"display_name"`
	Description      string    `json:"description"`
	LogoURL          string    `json:"logo_url"`
	Website          string    `json:"website"`
	DocumentationURL string    `json:"documentation_url"`
	BaseURL          string    `json:"base_url"`
	APIVersion       string    `json:"api_version"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
