package providers

import (
	"context"
	"io"
)

// Request represents a unified request structure for all AI providers
type Request struct {
	Model       string                 `json:"model"`
	Messages    []Message              `json:"messages"`
	MaxTokens   int                    `json:"max_tokens,omitempty"`
	Temperature float64                `json:"temperature,omitempty"`
	TopP        float64                `json:"top_p,omitempty"`
	Stream      bool                   `json:"stream,omitempty"`
	Stop        []string               `json:"stop,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response represents a unified response structure from all AI providers
type Response struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choice  `json:"choices"`
	Usage   Usage     `json:"usage"`
	Error   *APIError `json:"error,omitempty"`
}

// Choice represents a response choice
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// APIError represents an error response from the provider
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
}

// StreamResponse represents a streaming response chunk
type StreamResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []StreamChoice `json:"choices"`
}

// StreamChoice represents a streaming choice
type StreamChoice struct {
	Index        int         `json:"index"`
	Delta        StreamDelta `json:"delta"`
	FinishReason *string     `json:"finish_reason,omitempty"`
}

// StreamDelta represents a streaming delta
type StreamDelta struct {
	Content string `json:"content"`
}

// Provider defines the interface that all AI providers must implement
type Provider interface {
	// Name returns the provider name (e.g., "openai", "anthropic", etc.)
	Name() string

	// SendRequest sends a request to the provider and returns the response
	SendRequest(ctx context.Context, req *Request) (*Response, error)

	// SendStreamRequest sends a streaming request to the provider
	SendStreamRequest(ctx context.Context, req *Request) (io.ReadCloser, error)

	// GetModels returns the list of available models for this provider
	GetModels(ctx context.Context) ([]ModelInfo, error)

	// GetModelInfo returns detailed information about a specific model
	GetModelInfo(ctx context.Context, modelID string) (*ModelInfo, error)

	// ValidateModel checks if a model is valid for this provider
	ValidateModel(ctx context.Context, modelID string) error

	// GetPricing returns the pricing information for a model
	GetPricing(ctx context.Context, modelID string) (*Pricing, error)
}

// ModelInfo contains information about a model
type ModelInfo struct {
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	OwnedBy     string   `json:"owned_by"`
	Permission  []string `json:"permission"`
	MaxTokens   int      `json:"max_tokens"`
	ContextSize int      `json:"context_size"`
}

// Pricing contains pricing information for a model
type Pricing struct {
	InputCost  float64 `json:"input_cost"`  // Cost per 1K input tokens
	OutputCost float64 `json:"output_cost"` // Cost per 1K output tokens
	Currency   string  `json:"currency"`
}

// Config contains provider-specific configuration
type Config struct {
	APIKey     string            `json:"api_key"`
	BaseURL    string            `json:"base_url"`
	Headers    map[string]string `json:"headers"`
	Timeout    int               `json:"timeout"`
	MaxRetries int               `json:"max_retries"`
	RateLimit  RateLimitConfig   `json:"rate_limit"`
}

// RateLimitConfig contains rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int `json:"requests_per_minute"`
	TokensPerMinute   int `json:"tokens_per_minute"`
}

// ProviderFactory creates provider instances
type ProviderFactory interface {
	Create(config Config) (Provider, error)
}
