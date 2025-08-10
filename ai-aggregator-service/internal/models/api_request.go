package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// APIRequest represents the api_requests table
type APIRequest struct {
	bun.BaseModel `bun:"table:api_requests"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	APIKeyID       uuid.UUID  `bun:"api_key_id,notnull,type:uuid"`
	UserID         *uuid.UUID `bun:"user_id,type:uuid"`
	OrganizationID *uuid.UUID `bun:"organization_id,type:uuid"`
	ProviderID     uuid.UUID  `bun:"provider_id,notnull,type:uuid"`
	ModelID        uuid.UUID  `bun:"model_id,notnull,type:uuid"`
	RequestID      string     `bun:"request_id,notnull,type:varchar(255)"`
	RequestBody    JSONB      `bun:"request_body,type:jsonb,notnull"`
	RequestHeaders JSONB      `bun:"request_headers,type:jsonb,default:'{}'"`
	RequestMethod  string     `bun:"request_method,notnull,type:varchar(10)"`
	RequestURL     string     `bun:"request_url,notnull,type:text"`
	RequestSize    int        `bun:"request_size,type:integer,default:0"`
	Status         string     `bun:"status,notnull,type:varchar(50),default:'pending'"`
	StartedAt      time.Time  `bun:"started_at,notnull,default:current_timestamp"`
	CompletedAt    *time.Time `bun:"completed_at"`
	ErrorMessage   *string    `bun:"error_message,type:text"`
	ProcessingTime *int       `bun:"processing_time,type:integer"`

	// Relations
	APIKey             *APIKey             `bun:"rel:belongs-to,join:api_key_id=id"`
	User               *User               `bun:"rel:belongs-to,join:user_id=id"`
	Organization       *Organization       `bun:"rel:belongs-to,join:organization_id=id"`
	Provider           *Provider           `bun:"rel:belongs-to,join:provider_id=id"`
	Model              *Model              `bun:"rel:belongs-to,join:model_id=id"`
	APIResponse        *APIResponse        `bun:"rel:has-one,join:id=api_request_id"`
	BillingTransaction *BillingTransaction `bun:"rel:has-one,join:id=api_request_id"`
}

// Ensure APIRequest implements bun.BeforeAppendModelHook
var _ bun.BeforeAppendModelHook = (*APIRequest)(nil)

// BeforeAppendModel implements bun.BeforeAppendModelHook
func (m *APIRequest) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

// TableName returns the table name for APIRequest
func (APIRequest) TableName() string {
	return "api_requests"
}
