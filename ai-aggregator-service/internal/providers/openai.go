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

// OpenAIProvider implements the Provider interface for OpenAI
type OpenAIProvider struct {
	config Config
	client *http.Client
}

// NewOpenAIProvider creates a new OpenAI provider instance
func NewOpenAIProvider(config Config) *OpenAIProvider {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.openai.com/v1"
	}
	if config.Timeout == 0 {
		config.Timeout = 30
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	return &OpenAIProvider{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

// Name returns the provider name
func (p *OpenAIProvider) Name() string {
	return "openai"
}

// SendRequest sends a request to OpenAI
func (p *OpenAIProvider) SendRequest(ctx context.Context, req *Request) (*Response, error) {
	openaiReq := p.convertToOpenAIRequest(req)

	jsonData, err := json.Marshal(openaiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
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

	var openaiResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return p.convertFromOpenAIResponse(&openaiResp), nil
}

// SendStreamRequest sends a streaming request to OpenAI
func (p *OpenAIProvider) SendStreamRequest(ctx context.Context, req *Request) (io.ReadCloser, error) {
	openaiReq := p.convertToOpenAIRequest(req)
	openaiReq.Stream = true

	jsonData, err := json.Marshal(openaiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
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

// GetModels returns the list of available OpenAI models
func (p *OpenAIProvider) GetModels(ctx context.Context) ([]ModelInfo, error) {
	httpReq, err := http.NewRequestWithContext(ctx, "GET", p.config.BaseURL+"/models", nil)
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

	var modelsResp OpenAIModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var models []ModelInfo
	for _, model := range modelsResp.Data {
		models = append(models, ModelInfo{
			ID:      model.ID,
			Object:  model.Object,
			OwnedBy: model.OwnedBy,
		})
	}

	return models, nil
}

// GetModelInfo returns detailed information about a specific model
func (p *OpenAIProvider) GetModelInfo(ctx context.Context, modelID string) (*ModelInfo, error) {
	httpReq, err := http.NewRequestWithContext(ctx, "GET", p.config.BaseURL+"/models/"+modelID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	p.setHeaders(httpReq)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("model not found: %s", modelID)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var model OpenAIModel
	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &ModelInfo{
		ID:      model.ID,
		Object:  model.Object,
		OwnedBy: model.OwnedBy,
	}, nil
}

// ValidateModel checks if a model is valid for OpenAI
func (p *OpenAIProvider) ValidateModel(ctx context.Context, modelID string) error {
	_, err := p.GetModelInfo(ctx, modelID)
	return err
}

// GetPricing returns the pricing information for an OpenAI model
func (p *OpenAIProvider) GetPricing(ctx context.Context, modelID string) (*Pricing, error) {
	// OpenAI pricing data (as of 2024)
	pricingMap := map[string]Pricing{
		"gpt-4":             {InputCost: 0.03, OutputCost: 0.06, Currency: "usd"},
		"gpt-4-turbo":       {InputCost: 0.01, OutputCost: 0.03, Currency: "usd"},
		"gpt-3.5-turbo":     {InputCost: 0.0005, OutputCost: 0.0015, Currency: "usd"},
		"gpt-3.5-turbo-16k": {InputCost: 0.003, OutputCost: 0.004, Currency: "usd"},
		"text-davinci-003":  {InputCost: 0.02, OutputCost: 0.02, Currency: "usd"},
		"text-curie-001":    {InputCost: 0.002, OutputCost: 0.002, Currency: "usd"},
		"text-babbage-001":  {InputCost: 0.0005, OutputCost: 0.0005, Currency: "usd"},
		"text-ada-001":      {InputCost: 0.0004, OutputCost: 0.0004, Currency: "usd"},
		"dall-e-3":          {InputCost: 0.04, OutputCost: 0.08, Currency: "usd"},
		"whisper-1":         {InputCost: 0.006, OutputCost: 0.006, Currency: "usd"},
	}

	if pricing, exists := pricingMap[modelID]; exists {
		return &pricing, nil
	}

	// Default pricing for unknown models
	return &Pricing{
		InputCost:  0.01,
		OutputCost: 0.03,
		Currency:   "usd",
	}, nil
}

// setHeaders sets the required headers for OpenAI API
func (p *OpenAIProvider) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.config.APIKey)

	for key, value := range p.config.Headers {
		req.Header.Set(key, value)
	}
}

// convertToOpenAIRequest converts our unified request to OpenAI format
func (p *OpenAIProvider) convertToOpenAIRequest(req *Request) OpenAIRequest {
	return OpenAIRequest{
		Model:       req.Model,
		Messages:    p.convertMessages(req.Messages),
		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
		TopP:        req.TopP,
		Stream:      req.Stream,
		Stop:        req.Stop,
	}
}

// convertMessages converts our message format to OpenAI format
func (p *OpenAIProvider) convertMessages(messages []Message) []OpenAIMessage {
	var openaiMessages []OpenAIMessage
	for _, msg := range messages {
		openaiMessages = append(openaiMessages, OpenAIMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return openaiMessages
}

// convertFromOpenAIResponse converts OpenAI response to our unified format
func (p *OpenAIProvider) convertFromOpenAIResponse(resp *OpenAIResponse) *Response {
	var choices []Choice
	for _, choice := range resp.Choices {
		choices = append(choices, Choice{
			Index:        choice.Index,
			Message:      Message{Role: choice.Message.Role, Content: choice.Message.Content},
			FinishReason: choice.FinishReason,
		})
	}

	return &Response{
		ID:      resp.ID,
		Object:  resp.Object,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: choices,
		Usage: Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}
}

// OpenAIRequest represents the request format for OpenAI API
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	MaxTokens   int             `json:"max_tokens,omitempty"`
	Temperature float64         `json:"temperature,omitempty"`
	TopP        float64         `json:"top_p,omitempty"`
	Stream      bool            `json:"stream,omitempty"`
	Stop        []string        `json:"stop,omitempty"`
}

// OpenAIMessage represents a message in OpenAI format
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response format from OpenAI API
type OpenAIResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []OpenAIChoice `json:"choices"`
	Usage   OpenAIUsage    `json:"usage"`
}

// OpenAIChoice represents a choice in OpenAI response
type OpenAIChoice struct {
	Index        int           `json:"index"`
	Message      OpenAIMessage `json:"message"`
	FinishReason string        `json:"finish_reason"`
}

// OpenAIUsage represents usage information in OpenAI response
type OpenAIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// OpenAIModelsResponse represents the models response from OpenAI
type OpenAIModelsResponse struct {
	Object string        `json:"object"`
	Data   []OpenAIModel `json:"data"`
}

// OpenAIModel represents a model in OpenAI
type OpenAIModel struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}
