package models

import (
	"time"

	"github.com/google/uuid"
)

type APIRequest struct {
	ID             uuid.UUID `json:"id" db:"id"`
	APIKeyID       uuid.UUID `json:"api_key_id" db:"api_key_id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	ProviderID     uuid.UUID `json:"provider_id" db:"provider_id"`
	ModelID        uuid.UUID `json:"model_id" db:"model_id"`
	RequestID      string    `json:"request_id" db:"request_id"`
	Method         string    `json:"method" db:"method"`
	Endpoint       string    `json:"endpoint" db:"endpoint"`
	Headers        JSONB     `json:"headers" db:"headers"`
	RequestBody    JSONB     `json:"request_body" db:"request_body"`
	StatusCode     int       `json:"status_code" db:"status_code"`
	ErrorMessage   string    `json:"error_message" db:"error_message"`
	RequestTokens  int       `json:"request_tokens" db:"request_tokens"`
	ResponseTokens int       `json:"response_tokens" db:"response_tokens"`
	TotalTokens    int       `json:"total_tokens" db:"total_tokens"`
	Cost           float64   `json:"cost" db:"cost"`
	Duration       int64     `json:"duration" db:"duration"`
	IPAddress      string    `json:"ip_address" db:"ip_address"`
	UserAgent      string    `json:"user_agent" db:"user_agent"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type CreateAPIRequest struct {
	APIKeyID       uuid.UUID `json:"api_key_id" validate:"required"`
	UserID         uuid.UUID `json:"user_id" validate:"required"`
	OrganizationID uuid.UUID `json:"organization_id" validate:"required"`
	ProviderID     uuid.UUID `json:"provider_id" validate:"required"`
	ModelID        uuid.UUID `json:"model_id" validate:"required"`
	RequestID      string    `json:"request_id" validate:"required"`
	Method         string    `json:"method" validate:"required,oneof=GET POST PUT DELETE PATCH"`
	Endpoint       string    `json:"endpoint" validate:"required,url"`
	Headers        JSONB     `json:"headers,omitempty"`
	RequestBody    JSONB     `json:"request_body,omitempty"`
}

type UpdateAPIRequest struct {
	StatusCode     *int     `json:"status_code,omitempty"`
	ErrorMessage   *string  `json:"error_message,omitempty"`
	RequestTokens  *int     `json:"request_tokens,omitempty" validate:"omitempty,min=0"`
	ResponseTokens *int     `json:"response_tokens,omitempty" validate:"omitempty,min=0"`
	TotalTokens    *int     `json:"total_tokens,omitempty" validate:"omitempty,min=0"`
	Cost           *float64 `json:"cost,omitempty" validate:"omitempty,min=0"`
	Duration       *int64   `json:"duration,omitempty" validate:"omitempty,min=0"`
}

type APIRequestResponse struct {
	ID             uuid.UUID `json:"id"`
	APIKeyID       uuid.UUID `json:"api_key_id"`
	UserID         uuid.UUID `json:"user_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	ProviderID     uuid.UUID `json:"provider_id"`
	ModelID        uuid.UUID `json:"model_id"`
	RequestID      string    `json:"request_id"`
	Method         string    `json:"method"`
	Endpoint       string    `json:"endpoint"`
	StatusCode     int       `json:"status_code"`
	ErrorMessage   string    `json:"error_message"`
	RequestTokens  int       `json:"request_tokens"`
	ResponseTokens int       `json:"response_tokens"`
	TotalTokens    int       `json:"total_tokens"`
	Cost           float64   `json:"cost"`
	Duration       int64     `json:"duration"`
	IPAddress      string    `json:"ip_address"`
	UserAgent      string    `json:"user_agent"`
	CreatedAt      time.Time `json:"created_at"`
}

type APIRequestWithDetails struct {
	APIRequestResponse
	APIKey   APIKeyResponse   `json:"api_key"`
	User     UserResponse     `json:"user"`
	Provider ProviderResponse `json:"provider"`
	Model    ModelResponse    `json:"model"`
}
