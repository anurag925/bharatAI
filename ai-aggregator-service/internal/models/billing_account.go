package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// BillingAccount represents the billing_accounts table
type BillingAccount struct {
	bun.BaseModel `bun:"table:billing_accounts"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	OrganizationID *uuid.UUID `bun:"organization_id,type:uuid"`
	UserID         *uuid.UUID `bun:"user_id,type:uuid"`
	AccountType    string     `bun:"account_type,notnull,type:varchar(50)"`
	AccountName    string     `bun:"account_name,notnull,type:varchar(255)"`
	Balance        float64    `bun:"balance,notnull,type:numeric,default:0.0"`
	Currency       string     `bun:"currency,notnull,type:varchar(3),default:'USD'"`
	Status         string     `bun:"status,notnull,type:varchar(50),default:'active'"`
	BillingEmail   string     `bun:"billing_email,type:varchar(255)"`
	BillingAddress JSONB      `bun:"billing_address,type:jsonb,default:'{}'"`
	Metadata       JSONB      `bun:"metadata,type:jsonb,default:'{}'"`
	IsActive       bool       `bun:"is_active,notnull,default:true"`

	// Relations
	Organization        *Organization         `bun:"rel:belongs-to,join:organization_id=id"`
	User                *User                 `bun:"rel:belongs-to,join:user_id=id"`
	BillingTransactions []*BillingTransaction `bun:"rel:has-many,join:id=billing_account_id"`
}

// Ensure BillingAccount implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*BillingAccount)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *BillingAccount) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for BillingAccount
func (BillingAccount) TableName() string {
	return "billing_accounts"
}
