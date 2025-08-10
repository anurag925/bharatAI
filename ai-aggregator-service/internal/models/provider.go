package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Provider represents the providers table
type Provider struct {
	bun.BaseModel `bun:"table:providers"`

	ID                uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	Name              string    `bun:"name,notnull,unique,type:varchar(100)"`
	DisplayName       string    `bun:"display_name,notnull,type:varchar(255)"`
	BaseURL           string    `bun:"base_url,notnull,type:varchar(500)"`
	APIKeyRequired    bool      `bun:"api_key_required,notnull,default:true"`
	IsActive          bool      `bun:"is_active,notnull,default:true"`
	RateLimitRPM      int       `bun:"rate_limit_rpm,default:1000"`
	RateLimitTPM      int       `bun:"rate_limit_tpm,default:100000"`
	Config            JSONB     `bun:"config,type:jsonb,default:'{}'"`
	SupportedFeatures JSONB     `bun:"supported_features,type:jsonb,default:'[]'"`

	// Relations
	Models       []*Model       `bun:"rel:has-many,join:id=provider_id"`
	APIRequests  []*APIRequest  `bun:"rel:has-many,join:id=provider_id"`
	APIResponses []*APIResponse `bun:"rel:has-many,join:id=provider_id"`
	RateLimits   []*RateLimit   `bun:"rel:has-many,join:id=provider_id"`
}

// TableName returns the table name for Provider
func (Provider) TableName() string {
	return "providers"
}

// Ensure Provider implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*Provider)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (p *Provider) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		p.CreatedAt = time.Now()
		p.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		p.UpdatedAt = time.Now()
	}
	return nil
}
