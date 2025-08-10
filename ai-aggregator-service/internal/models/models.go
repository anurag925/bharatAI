package models

import (
	"time"

	"github.com/google/uuid"
)

// ModelList contains all models for database initialization
var ModelList = []interface{}{
	(*Organization)(nil),
	(*Provider)(nil),
	(*User)(nil),
	(*APIKey)(nil),
	(*Model)(nil),
	(*APIRequest)(nil),
	(*APIResponse)(nil),
	(*BillingAccount)(nil),
	(*BillingTransaction)(nil),
	(*RateLimit)(nil),
}

// NullUUID returns a nil UUID pointer
func NullUUID() *uuid.UUID {
	return nil
}

// UUIDPtr returns a pointer to the given UUID
func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

// StringPtr returns a pointer to the given string
func StringPtr(s string) *string {
	return &s
}

// TimePtr returns a pointer to the given time
func TimePtr(t time.Time) *time.Time {
	return &t
}

// IntPtr returns a pointer to the given int
func IntPtr(i int) *int {
	return &i
}

// Float64Ptr returns a pointer to the given float64
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr returns a pointer to the given bool
func BoolPtr(b bool) *bool {
	return &b
}
