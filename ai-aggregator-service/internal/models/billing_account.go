package models

import (
	"time"

	"github.com/google/uuid"
)

type BillingAccount struct {
	ID               uuid.UUID `json:"id" db:"id"`
	OrganizationID   uuid.UUID `json:"organization_id" db:"organization_id"`
	StripeCustomerID string    `json:"stripe_customer_id" db:"stripe_customer_id"`
	Balance          float64   `json:"balance" db:"balance"`
	Currency         string    `json:"currency" db:"currency"`
	Status           string    `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBillingAccount struct {
	OrganizationID   uuid.UUID `json:"organization_id" validate:"required"`
	StripeCustomerID string    `json:"stripe_customer_id" validate:"required"`
	Currency         string    `json:"currency" validate:"required,oneof=usd eur gbp jpy"`
}

type UpdateBillingAccount struct {
	StripeCustomerID *string  `json:"stripe_customer_id,omitempty"`
	Balance          *float64 `json:"balance,omitempty"`
	Currency         *string  `json:"currency,omitempty" validate:"omitempty,oneof=usd eur gbp jpy"`
	Status           *string  `json:"status,omitempty" validate:"omitempty,oneof=active suspended closed"`
}

type BillingAccountResponse struct {
	ID               uuid.UUID `json:"id"`
	OrganizationID   uuid.UUID `json:"organization_id"`
	StripeCustomerID string    `json:"stripe_customer_id"`
	Balance          float64   `json:"balance"`
	Currency         string    `json:"currency"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type BillingAccountWithOrganization struct {
	BillingAccountResponse
	Organization OrganizationResponse `json:"organization"`
}
