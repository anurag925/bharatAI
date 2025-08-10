package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Organization represents the organizations table
type Organization struct {
	bun.BaseModel `bun:"table:organizations"`

	ID           uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	Name         string    `bun:"name,notnull,type:varchar(255)"`
	Slug         string    `bun:"slug,notnull,unique,type:varchar(100)"`
	PlanType     string    `bun:"plan_type,notnull,default:'free',type:varchar(50)"`
	BillingEmail string    `bun:"billing_email,type:varchar(255)"`
	Metadata     JSONB     `bun:"metadata,type:jsonb,default:'{}'"`
	IsActive     bool      `bun:"is_active,notnull,default:true"`

	// Relations
	Users           []*User           `bun:"rel:has-many,join:id=organization_id"`
	APIKeys         []*APIKey         `bun:"rel:has-many,join:id=organization_id"`
	BillingAccounts []*BillingAccount `bun:"rel:has-many,join:id=organization_id"`
	RateLimits      []*RateLimit      `bun:"rel:has-many,join:id=organization_id"`
}

// TableName returns the table name for Organization
func (Organization) TableName() string {
	return "organizations"
}

// Ensure Organization implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*Organization)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (o *Organization) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		o.CreatedAt = time.Now()
		o.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		o.UpdatedAt = time.Now()
	}
	return nil
}
