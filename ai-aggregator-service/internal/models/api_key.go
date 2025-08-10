package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// APIKey represents the api_keys table
type APIKey struct {
	bun.BaseModel `bun:"table:api_keys"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	KeyHash        string     `bun:"key_hash,notnull,unique,type:varchar(255)"`
	UserID         *uuid.UUID `bun:"user_id,type:uuid"`
	OrganizationID *uuid.UUID `bun:"organization_id,type:uuid"`
	Name           string     `bun:"name,notnull,type:varchar(255)"`
	Permissions    JSONB      `bun:"permissions,type:jsonb,default:'[]'"`
	IsActive       bool       `bun:"is_active,notnull,default:true"`
	LastUsedAt     *time.Time `bun:"last_used_at"`
	ExpiresAt      *time.Time `bun:"expires_at"`

	// Relations
	User         *User         `bun:"rel:belongs-to,join:user_id=id"`
	Organization *Organization `bun:"rel:belongs-to,join:organization_id=id"`
	APIRequests  []*APIRequest `bun:"rel:has-many,join:id=api_key_id"`
	RateLimits   []*RateLimit  `bun:"rel:has-many,join:id=api_key_id"`
}

// Ensure APIKey implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*APIKey)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *APIKey) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for APIKey
func (APIKey) TableName() string {
	return "api_keys"
}
