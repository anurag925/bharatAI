package models

import (
	"time"

	"github.com/google/uuid"
)

type BillingTransaction struct {
	ID               uuid.UUID `json:"id" db:"id"`
	BillingAccountID uuid.UUID `json:"billing_account_id" db:"billing_account_id"`
	APIRequestID     uuid.UUID `json:"api_request_id" db:"api_request_id"`
	Type             string    `json:"type" db:"type"`
	Amount           float64   `json:"amount" db:"amount"`
	Currency         string    `json:"currency" db:"currency"`
	Description      string    `json:"description" db:"description"`
	Metadata         JSONB     `json:"metadata" db:"metadata"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type CreateBillingTransaction struct {
	BillingAccountID uuid.UUID `json:"billing_account_id" validate:"required"`
	APIRequestID     uuid.UUID `json:"api_request_id" validate:"required"`
	Type             string    `json:"type" validate:"required,oneof=charge credit refund"`
	Amount           float64   `json:"amount" validate:"required"`
	Currency         string    `json:"currency" validate:"required,oneof=usd eur gbp jpy"`
	Description      string    `json:"description" validate:"required"`
	Metadata         JSONB     `json:"metadata,omitempty"`
}

type BillingTransactionResponse struct {
	ID               uuid.UUID `json:"id"`
	BillingAccountID uuid.UUID `json:"billing_account_id"`
	APIRequestID     uuid.UUID `json:"api_request_id"`
	Type             string    `json:"type"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Description      string    `json:"description"`
	Metadata         JSONB     `json:"metadata"`
	CreatedAt        time.Time `json:"created_at"`
}

type BillingTransactionWithDetails struct {
	BillingTransactionResponse
	BillingAccount BillingAccountResponse `json:"billing_account"`
	APIRequest     APIRequestResponse     `json:"api_request"`
}
