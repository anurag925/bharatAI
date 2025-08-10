package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// RateLimit represents the rate_limits table
type RateLimit struct {
	bun.BaseModel `bun:"table:rate_limits"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	APIKeyID       *uuid.UUID `bun:"api_key_id,type:uuid"`
	UserID         *uuid.UUID `bun:"user_id,type:uuid"`
	OrganizationID *uuid.UUID `bun:"organization_id,type:uuid"`
	ProviderID     *uuid.UUID `bun:"provider_id,type:uuid"`
	ModelID        *uuid.UUID `bun:"model_id,type:uuid"`
	LimitType      string     `bun:"limit_type,notnull,type:varchar(50)"`
	LimitValue     int        `bun:"limit_value,notnull,type:integer"`
	CurrentUsage   int        `bun:"current_usage,notnull,type:integer,default:0"`
	WindowSize     string     `bun:"window_size,type:varchar(50)"`
	WindowStart    time.Time  `bun:"window_start,notnull,default:current_timestamp"`
	WindowEnd      time.Time  `bun:"window_end,notnull,default:current_timestamp"`
	IsActive       bool       `bun:"is_active,notnull,default:true"`
	Metadata       JSONB      `bun:"metadata,type:jsonb,default:'{}'"`

	// Relations
	APIKey       *APIKey       `bun:"rel:belongs-to,join:api_key_id=id"`
	User         *User         `bun:"rel:belongs-to,join:user_id=id"`
	Organization *Organization `bun:"rel:belongs-to,join:organization_id=id"`
	Provider     *Provider     `bun:"rel:belongs-to,join:provider_id=id"`
	Model        *Model        `bun:"rel:belongs-to,join:model_id=id"`
}

// Ensure RateLimit implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*RateLimit)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *RateLimit) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for RateLimit
func (RateLimit) TableName() string {
	return "rate_limits"
}
