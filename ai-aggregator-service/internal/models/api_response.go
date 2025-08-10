package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// APIResponse represents the api_responses table
type APIResponse struct {
	bun.BaseModel `bun:"table:api_responses"`

	ID              uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	APIRequestID    uuid.UUID `bun:"api_request_id,notnull,type:uuid"`
	ProviderID      uuid.UUID `bun:"provider_id,notnull,type:uuid"`
	ModelID         uuid.UUID `bun:"model_id,notnull,type:uuid"`
	ResponseBody    JSONB     `bun:"response_body,type:jsonb,notnull"`
	ResponseHeaders JSONB     `bun:"response_headers,type:jsonb,default:'{}'"`
	StatusCode      int       `bun:"status_code,notnull,type:integer"`
	ResponseSize    int       `bun:"response_size,type:integer,default:0"`
	ProcessingTime  int       `bun:"processing_time,type:integer"`
	ErrorCode       *string   `bun:"error_code,type:varchar(50)"`
	ErrorMessage    *string   `bun:"error_message,type:text"`
	TokenUsage      JSONB     `bun:"token_usage,type:jsonb,default:'{}'"`
	Cost            float64   `bun:"cost,type:numeric,default:0.0"`

	// Relations
	APIRequest *APIRequest `bun:"rel:belongs-to,join:api_request_id=id"`
	Provider   *Provider   `bun:"rel:belongs-to,join:provider_id=id"`
	Model      *Model      `bun:"rel:belongs-to,join:model_id=id"`
}

// Ensure APIResponse implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*APIResponse)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *APIResponse) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for APIResponse
func (APIResponse) TableName() string {
	return "api_responses"
}
