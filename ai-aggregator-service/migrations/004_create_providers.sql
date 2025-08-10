-- Create providers table
CREATE TABLE IF NOT EXISTS providers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    base_url VARCHAR(500) NOT NULL,
    api_key_required BOOLEAN DEFAULT TRUE,
    is_active BOOLEAN DEFAULT TRUE,
    rate_limit_rpm INTEGER DEFAULT 1000,
    rate_limit_tpm INTEGER DEFAULT 100000,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    config JSONB DEFAULT '{}'::jsonb,
    supported_features JSONB DEFAULT '[]'::jsonb
);

-- Create indexes for providers
CREATE INDEX IF NOT EXISTS idx_providers_name ON providers(name);
CREATE INDEX IF NOT EXISTS idx_providers_is_active ON providers(is_active);
CREATE INDEX IF NOT EXISTS idx_providers_created_at ON providers(created_at);

-- Insert default providers
INSERT INTO providers (name, display_name, base_url, api_key_required, rate_limit_rpm, rate_limit_tpm, supported_features) VALUES
    ('openai', 'OpenAI', 'https://api.openai.com/v1', true, 10000, 1000000, '["chat", "completions", "embeddings", "images", "audio", "files"]'::jsonb),
    ('anthropic', 'Anthropic', 'https://api.anthropic.com', true, 1000, 40000, '["chat", "completions"]'::jsonb),
    ('google', 'Google AI', 'https://generativelanguage.googleapis.com', true, 3000, 60000, '["chat", "completions", "embeddings"]'::jsonb),
    ('cohere', 'Cohere', 'https://api.cohere.ai', true, 10000, 1000000, '["chat", "completions", "embeddings", "classify"]'::jsonb),
    ('mistral', 'Mistral AI', 'https://api.mistral.ai', true, 2000, 200000, '["chat", "completions", "embeddings"]'::jsonb)
ON CONFLICT (name) DO NOTHING;