package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config holds all configuration for the application
type Config struct {
	Server    ServerConfig    `envPrefix:"SERVER_"`
	Database  DatabaseConfig  `envPrefix:"DATABASE_"`
	Redis     RedisConfig     `envPrefix:"REDIS_"`
	NATS      NATSConfig      `envPrefix:"NATS_"`
	Logging   LoggingConfig   `envPrefix:"LOGGING_"`
	Auth      AuthConfig      `envPrefix:"AUTH_"`
	Metrics   MetricsConfig   `envPrefix:"METRICS_"`
	Providers ProvidersConfig `envPrefix:"PROVIDERS_"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Host         string        `env:"HOST" envDefault:"0.0.0.0"`
	Port         int           `env:"PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"30s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"30s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" envDefault:"120s"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	DBName   string `env:"DBNAME" envDefault:"ai_aggregator_db"`
	SSLMode  string `env:"SSL_MODE" envDefault:"disable"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"6379"`
	Password string `env:"PASSWORD" envDefault:""`
	DB       int    `env:"DB" envDefault:"0"`
}

// NATSConfig holds NATS configuration
type NATSConfig struct {
	URL string `env:"URL" envDefault:"nats://localhost:4222"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `env:"LEVEL" envDefault:"info"`
	Format string `env:"FORMAT" envDefault:"json"`
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	JWTSecret     string        `env:"JWT_SECRET" envDefault:"your-secret-key-change-this-in-production"`
	JWTExpiration time.Duration `env:"JWT_EXPIRATION" envDefault:"24h"`
}

// MetricsConfig holds metrics configuration
type MetricsConfig struct {
	Enabled bool   `env:"ENABLED" envDefault:"true"`
	Port    int    `env:"PORT" envDefault:"9090"`
	Path    string `env:"PATH" envDefault:"/metrics"`
}

// ProviderConfig holds configuration for AI providers
type ProviderConfig struct {
	Name    string            `env:"NAME"`
	APIKey  string            `env:"API_KEY"`
	BaseURL string            `env:"BASE_URL"`
	Models  []string          `env:"MODELS"`
	Headers map[string]string `env:"HEADERS"`
}

// ProvidersConfig holds all provider configurations
type ProvidersConfig struct {
	OpenAI    ProviderConfig `envPrefix:"OPENAI_"`
	Anthropic ProviderConfig `envPrefix:"ANTHROPIC_"`
	GoogleAI  ProviderConfig `envPrefix:"GOOGLE_AI_"`
	Cohere    ProviderConfig `envPrefix:"COHERE_"`
}

// Provider-specific configs with defaults
type openAIConfig struct {
	Name   string `env:"NAME" envDefault:"openai"`
	APIKey string `env:"API_KEY"`
}

type anthropicConfig struct {
	Name   string `env:"NAME" envDefault:"anthropic"`
	APIKey string `env:"API_KEY"`
}

type googleAIConfig struct {
	Name   string `env:"NAME" envDefault:"google_ai"`
	APIKey string `env:"API_KEY"`
}

type cohereConfig struct {
	Name   string `env:"NAME" envDefault:"cohere"`
	APIKey string `env:"API_KEY"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// First load the main config with AGG_ prefix
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// Load provider configurations separately to handle different prefixes
	// For standard provider env vars (without AGG_ prefix)
	openAI := openAIConfig{}
	anthropic := anthropicConfig{}
	googleAI := googleAIConfig{}
	cohere := cohereConfig{}

	if err := env.Parse(&openAI); err != nil {
		return nil, err
	}
	if err := env.Parse(&anthropic); err != nil {
		return nil, err
	}
	if err := env.Parse(&googleAI); err != nil {
		return nil, err
	}
	if err := env.Parse(&cohere); err != nil {
		return nil, err
	}

	// Set provider configs
	cfg.Providers.OpenAI = ProviderConfig{
		Name:   openAI.Name,
		APIKey: openAI.APIKey,
	}
	cfg.Providers.Anthropic = ProviderConfig{
		Name:   anthropic.Name,
		APIKey: anthropic.APIKey,
	}
	cfg.Providers.GoogleAI = ProviderConfig{
		Name:   googleAI.Name,
		APIKey: googleAI.APIKey,
	}
	cfg.Providers.Cohere = ProviderConfig{
		Name:   cohere.Name,
		APIKey: cohere.APIKey,
	}

	return cfg, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return &ConfigError{Field: "server.port", Value: c.Server.Port, Message: "port must be between 1 and 65535"}
	}

	if c.Database.Host == "" {
		return &ConfigError{Field: "database.host", Value: "", Message: "database host is required"}
	}

	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return &ConfigError{Field: "database.port", Value: c.Database.Port, Message: "port must be between 1 and 65535"}
	}

	if c.Redis.Port <= 0 || c.Redis.Port > 65535 {
		return &ConfigError{Field: "redis.port", Value: c.Redis.Port, Message: "port must be between 1 and 65535"}
	}

	return nil
}

// ConfigError represents a configuration validation error
type ConfigError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}

// SetupLogger configures the global logger
func SetupLogger(config *LoggingConfig) {
	var level slog.Level
	switch config.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler
	switch config.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
