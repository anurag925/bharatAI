package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// User represents the users table
type User struct {
	bun.BaseModel `bun:"table:users"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	OrganizationID uuid.UUID  `bun:"organization_id,type:uuid,notnull"`
	Email          string     `bun:"email,notnull,unique,type:varchar(255)"`
	FullName       string     `bun:"full_name,type:varchar(255)"`
	Role           string     `bun:"role,notnull,default:'member',type:varchar(50)"`
	IsActive       bool       `bun:"is_active,notnull,default:true"`
	LastLoginAt    *time.Time `bun:"last_login_at"`
	Metadata       JSONB      `bun:"metadata,type:jsonb,default:'{}'"`

	// Relations
	Organization    *Organization     `bun:"rel:belongs-to,join:organization_id=id"`
	APIKeys         []*APIKey         `bun:"rel:has-many,join:id=user_id"`
	BillingAccounts []*BillingAccount `bun:"rel:has-many,join:id=user_id"`
	APIRequests     []*APIRequest     `bun:"rel:has-many,join:id=user_id"`
	RateLimits      []*RateLimit      `bun:"rel:has-many,join:id=user_id"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}

// Ensure User implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*User)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}
