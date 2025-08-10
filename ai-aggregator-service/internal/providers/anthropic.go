package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// AnthropicProvider implements the Provider interface for Anthropic
type AnthropicProvider struct {
	config Config
	client *http.Client
}

// NewAnthropicProvider creates a new Anthropic provider instance
func NewAnthropicProvider(config Config) *AnthropicProvider {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.anthropic.com/v1"
	}
	if config.Timeout == 0 {
		config.Timeout = 30
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	return &AnthropicProvider{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

// Name returns the provider name
func (p *AnthropicProvider) Name() string {
	return "anthropic"
}

// SendRequest sends a request to Anthropic
func (p *AnthropicProvider) SendRequest(ctx context.Context, req *Request) (*Response, error) {
	anthropicReq := p.convertToAnthropicRequest(req)

	jsonData, err := json.Marshal(anthropicReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.config.BaseURL+"/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	p.setHeaders(httpReq)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var anthropicResp AnthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&anthropicResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return p.convertFromAnthropicResponse(&anthropicResp, req.Model), nil
}

// SendStreamRequest sends a streaming request to Anthropic
func (p *AnthropicProvider) SendStreamRequest(ctx context.Context, req *Request) (io.ReadCloser, error) {
	anthropicReq := p.convertToAnthropicRequest(req)
	anthropicReq.Stream = true

	jsonData, err := json.Marshal(anthropicReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.config.BaseURL+"/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	p.setHeaders(httpReq)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return resp.Body, nil
}

// GetModels returns the list of available Anthropic models
func (p *AnthropicProvider) GetModels(ctx context.Context) ([]ModelInfo, error) {
	// Anthropic doesn't have a models endpoint, so we return hardcoded models
	models := []ModelInfo{
		{
			ID:          "claude-3-5-sonnet-20241022",
			Object:      "model",
			OwnedBy:     "anthropic",
			MaxTokens:   4096,
			ContextSize: 200000,
		},
		{
			ID:          "claude-3-5-haiku-20241022",
			Object:      "model",
			OwnedBy:     "anthropic",
			MaxTokens:   4096,
			ContextSize: 200000,
		},
		{
			ID:          "claude-3-opus-20240229",
			Object:      "model",
			OwnedBy:     "anthropic",
			MaxTokens:   4096,
			ContextSize: 200000,
		},
	}

	return models, nil
}

// GetModelInfo returns detailed information about a specific model
func (p *AnthropicProvider) GetModelInfo(ctx context.Context, modelID string) (*ModelInfo, error) {
	models, err := p.GetModels(ctx)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		if model.ID == modelID {
			return &model, nil
		}
	}

	return nil, fmt.Errorf("model not found: %s", modelID)
}

// ValidateModel checks if a model is valid for Anthropic
func (p *AnthropicProvider) ValidateModel(ctx context.Context, modelID string) error {
	_, err := p.GetModelInfo(ctx, modelID)
	return err
}

// GetPricing returns the pricing information for an Anthropic model
func (p *AnthropicProvider) GetPricing(ctx context.Context, modelID string) (*Pricing, error) {
	// Anthropic pricing data (as of 2024)
	pricingMap := map[string]Pricing{
		"claude-3-5-sonnet-20241022": {InputCost: 0.003, OutputCost: 0.015, Currency: "usd"},
		"claude-3-5-haiku-20241022":  {InputCost: 0.0008, OutputCost: 0.004, Currency: "usd"},
		"claude-3-opus-20240229":     {InputCost: 0.015, OutputCost: 0.075, Currency: "usd"},
	}

	if pricing, exists := pricingMap[modelID]; exists {
		return &pricing, nil
	}

	// Default pricing for unknown models
	return &Pricing{
		InputCost:  0.003,
		OutputCost: 0.015,
		Currency:   "usd",
	}, nil
}

// setHeaders sets the required headers for Anthropic API
func (p *AnthropicProvider) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.config.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	for key, value := range p.config.Headers {
		req.Header.Set(key, value)
	}
}

// convertToAnthropicRequest converts our unified request to Anthropic format
func (p *AnthropicProvider) convertToAnthropicRequest(req *Request) AnthropicRequest {
	return AnthropicRequest{
		Model:         req.Model,
		MaxTokens:     req.MaxTokens,
		Temperature:   req.Temperature,
		TopP:          req.TopP,
		Stream:        req.Stream,
		StopSequences: req.Stop,
		Messages:      p.convertMessages(req.Messages),
	}
}

// convertMessages converts our message format to Anthropic format
func (p *AnthropicProvider) convertMessages(messages []Message) []AnthropicMessage {
	var anthropicMessages []AnthropicMessage
	for _, msg := range messages {
		anthropicMessages = append(anthropicMessages, AnthropicMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return anthropicMessages
}

// convertFromAnthropicResponse converts Anthropic response to our unified format
func (p *AnthropicProvider) convertFromAnthropicResponse(resp *AnthropicResponse, model string) *Response {
	var choices []Choice
	for _, content := range resp.Content {
		if content.Type == "text" {
			choices = append(choices, Choice{
				Index: 0,
				Message: Message{
					Role:    "assistant",
					Content: content.Text,
				},
				FinishReason: resp.StopReason,
			})
		}
	}

	return &Response{
		ID:      resp.ID,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: choices,
		Usage: Usage{
			PromptTokens:     resp.Usage.InputTokens,
			CompletionTokens: resp.Usage.OutputTokens,
			TotalTokens:      resp.Usage.InputTokens + resp.Usage.OutputTokens,
		},
	}
}

// AnthropicRequest represents the request format for Anthropic API
type AnthropicRequest struct {
	Model         string             `json:"model"`
	MaxTokens     int                `json:"max_tokens,omitempty"`
	Temperature   float64            `json:"temperature,omitempty"`
	TopP          float64            `json:"top_p,omitempty"`
	Stream        bool               `json:"stream,omitempty"`
	StopSequences []string           `json:"stop_sequences,omitempty"`
	Messages      []AnthropicMessage `json:"messages"`
}

// AnthropicMessage represents a message in Anthropic format
type AnthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AnthropicResponse represents the response format from Anthropic API
type AnthropicResponse struct {
	ID           string             `json:"id"`
	Type         string             `json:"type"`
	Role         string             `json:"role"`
	Content      []AnthropicContent `json:"content"`
	Model        string             `json:"model"`
	StopReason   string             `json:"stop_reason"`
	StopSequence *string            `json:"stop_sequence,omitempty"`
	Usage        AnthropicUsage     `json:"usage"`
}

// AnthropicContent represents content in Anthropic response
type AnthropicContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// AnthropicUsage represents usage information in Anthropic response
type AnthropicUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}
