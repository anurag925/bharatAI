package models

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID            uuid.UUID `json:"id" db:"id"`
	ProviderID    uuid.UUID `json:"provider_id" db:"provider_id"`
	Name          string    `json:"name" db:"name"`
	DisplayName   string    `json:"display_name" db:"display_name"`
	Description   string    `json:"description" db:"description"`
	ModelType     string    `json:"model_type" db:"model_type"`
	Capabilities  JSONB     `json:"capabilities" db:"capabilities"`
	ContextWindow int       `json:"context_window" db:"context_window"`
	MaxTokens     int       `json:"max_tokens" db:"max_tokens"`
	InputPrice    float64   `json:"input_price" db:"input_price"`
	OutputPrice   float64   `json:"output_price" db:"output_price"`
	RequestPrice  float64   `json:"request_price" db:"request_price"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	Config        JSONB     `json:"config" db:"config"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type CreateModel struct {
	ProviderID    uuid.UUID `json:"provider_id" validate:"required"`
	Name          string    `json:"name" validate:"required,min=2,max=100"`
	DisplayName   string    `json:"display_name" validate:"required,min=2,max=100"`
	Description   string    `json:"description" validate:"max=500"`
	ModelType     string    `json:"model_type" validate:"required,oneof=text chat image audio video embedding"`
	Capabilities  JSONB     `json:"capabilities,omitempty"`
	ContextWindow int       `json:"context_window" validate:"min=0"`
	MaxTokens     int       `json:"max_tokens" validate:"min=0"`
	InputPrice    float64   `json:"input_price" validate:"min=0"`
	OutputPrice   float64   `json:"output_price" validate:"min=0"`
	RequestPrice  float64   `json:"request_price" validate:"min=0"`
	Config        JSONB     `json:"config,omitempty"`
}

type UpdateModel struct {
	DisplayName   *string  `json:"display_name,omitempty" validate:"omitempty,min=2,max=100"`
	Description   *string  `json:"description,omitempty" validate:"omitempty,max=500"`
	ContextWindow *int     `json:"context_window,omitempty" validate:"omitempty,min=0"`
	MaxTokens     *int     `json:"max_tokens,omitempty" validate:"omitempty,min=0"`
	InputPrice    *float64 `json:"input_price,omitempty" validate:"omitempty,min=0"`
	OutputPrice   *float64 `json:"output_price,omitempty" validate:"omitempty,min=0"`
	RequestPrice  *float64 `json:"request_price,omitempty" validate:"omitempty,min=0"`
	IsActive      *bool    `json:"is_active,omitempty"`
	Config        JSONB    `json:"config,omitempty"`
}

type ModelResponse struct {
	ID            uuid.UUID `json:"id"`
	ProviderID    uuid.UUID `json:"provider_id"`
	Name          string    `json:"name"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	ModelType     string    `json:"model_type"`
	ContextWindow int       `json:"context_window"`
	MaxTokens     int       `json:"max_tokens"`
	InputPrice    float64   `json:"input_price"`
	OutputPrice   float64   `json:"output_price"`
	RequestPrice  float64   `json:"request_price"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ModelWithProvider struct {
	ModelResponse
	Provider ProviderResponse `json:"provider"`
}
