package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Model represents the models table
type Model struct {
	bun.BaseModel `bun:"table:models"`

	ID            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	ProviderID    uuid.UUID `bun:"provider_id,notnull,type:uuid"`
	Name          string    `bun:"name,notnull,type:varchar(255)"`
	DisplayName   string    `bun:"display_name,type:varchar(255)"`
	Description   string    `bun:"description,type:text"`
	ModelType     string    `bun:"model_type,notnull,type:varchar(50)"`
	IsActive      bool      `bun:"is_active,notnull,default:true"`
	Config        JSONB     `bun:"config,type:jsonb,default:'{}'"`
	Pricing       JSONB     `bun:"pricing,type:jsonb,default:'{}'"`
	ContextWindow int       `bun:"context_window,type:integer"`
	MaxTokens     int       `bun:"max_tokens,type:integer"`
	Temperature   float64   `bun:"temperature,type:numeric"`
	TopP          float64   `bun:"top_p,type:numeric"`

	// Relations
	Provider     *Provider      `bun:"rel:belongs-to,join:provider_id=id"`
	APIRequests  []*APIRequest  `bun:"rel:has-many,join:id=model_id"`
	APIResponses []*APIResponse `bun:"rel:has-many,join:id=model_id"`
	RateLimits   []*RateLimit   `bun:"rel:has-many,join:id=model_id"`
}

// Ensure Model implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*Model)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for Model
func (Model) TableName() string {
	return "models"
}
