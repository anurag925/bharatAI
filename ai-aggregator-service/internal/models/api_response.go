package models

import (
	"time"

	"github.com/google/uuid"
)

type APIResponse struct {
	ID           uuid.UUID `json:"id" db:"id"`
	RequestID    uuid.UUID `json:"request_id" db:"request_id"`
	Headers      JSONB     `json:"headers" db:"headers"`
	ResponseBody JSONB     `json:"response_body" db:"response_body"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateAPIResponse struct {
	RequestID    uuid.UUID `json:"request_id" validate:"required"`
	Headers      JSONB     `json:"headers,omitempty"`
	ResponseBody JSONB     `json:"response_body,omitempty"`
}

type APIResponseResponse struct {
	ID           uuid.UUID `json:"id"`
	RequestID    uuid.UUID `json:"request_id"`
	Headers      JSONB     `json:"headers"`
	ResponseBody JSONB     `json:"response_body"`
	CreatedAt    time.Time `json:"created_at"`
}

type APIResponseWithRequest struct {
	APIResponseResponse
	Request APIRequestResponse `json:"request"`
}
