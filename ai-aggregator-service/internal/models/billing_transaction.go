package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// BillingTransaction represents the billing_transactions table
type BillingTransaction struct {
	bun.BaseModel `bun:"table:billing_transactions"`

	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt        time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt        time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	BillingAccountID uuid.UUID  `bun:"billing_account_id,notnull,type:uuid"`
	APIRequestID     *uuid.UUID `bun:"api_request_id,type:uuid"`
	TransactionType  string     `bun:"transaction_type,notnull,type:varchar(50)"`
	Amount           float64    `bun:"amount,notnull,type:numeric"`
	Currency         string     `bun:"currency,notnull,type:varchar(3),default:'USD'"`
	Description      string     `bun:"description,type:text"`
	Metadata         JSONB      `bun:"metadata,type:jsonb,default:'{}'"`
	Status           string     `bun:"status,notnull,type:varchar(50),default:'completed'"`
	ReferenceID      *string    `bun:"reference_id,type:varchar(255)"`
	ProcessedAt      *time.Time `bun:"processed_at"`

	// Relations
	BillingAccount *BillingAccount `bun:"rel:belongs-to,join:billing_account_id=id"`
	APIRequest     *APIRequest     `bun:"rel:belongs-to,join:api_request_id=id"`
}

// Ensure BillingTransaction implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*BillingTransaction)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *BillingTransaction) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for BillingTransaction
func (BillingTransaction) TableName() string {
	return "billing_transactions"
}
