-- Create models table
CREATE TABLE IF NOT EXISTS models (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    provider_id UUID REFERENCES providers(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    model_type VARCHAR(50) NOT NULL DEFAULT 'chat',
    is_active BOOLEAN DEFAULT TRUE,
    context_window INTEGER,
    max_tokens INTEGER,
    input_cost_per_1k_tokens DECIMAL(10,6) DEFAULT 0.0,
    output_cost_per_1k_tokens DECIMAL(10,6) DEFAULT 0.0,
    config JSONB DEFAULT '{}'::jsonb,
    capabilities JSONB DEFAULT '[]'::jsonb,
    UNIQUE(provider_id, name)
);

-- Create indexes for models
CREATE INDEX IF NOT EXISTS idx_models_provider_id ON models(provider_id);
CREATE INDEX IF NOT EXISTS idx_models_name ON models(name);
CREATE INDEX IF NOT EXISTS idx_models_type ON models(model_type);
CREATE INDEX IF NOT EXISTS idx_models_is_active ON models(is_active);
CREATE INDEX IF NOT EXISTS idx_models_created_at ON models(created_at);

-- Insert default models
INSERT INTO models (provider_id, name, display_name, model_type, context_window, max_tokens, input_cost_per_1k_tokens, output_cost_per_1k_tokens, capabilities) VALUES
    ((SELECT id FROM providers WHERE name = 'openai'), 'gpt-4o', 'GPT-4o', 'chat', 128000, 4096, 0.005, 0.015, '["text", "vision", "code", "reasoning"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'openai'), 'gpt-4o-mini', 'GPT-4o Mini', 'chat', 128000, 16384, 0.00015, 0.0006, '["text", "vision", "code"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'openai'), 'gpt-3.5-turbo', 'GPT-3.5 Turbo', 'chat', 16385, 4096, 0.0005, 0.0015, '["text", "code"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'anthropic'), 'claude-3-5-sonnet-20241022', 'Claude 3.5 Sonnet', 'chat', 200000, 8192, 0.003, 0.015, '["text", "vision", "code", "reasoning"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'anthropic'), 'claude-3-5-haiku-20241022', 'Claude 3.5 Haiku', 'chat', 200000, 8192, 0.0008, 0.004, '["text", "vision", "code"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'google'), 'gemini-1.5-pro', 'Gemini 1.5 Pro', 'chat', 2097152, 8192, 0.00125, 0.005, '["text", "vision", "code", "reasoning"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'google'), 'gemini-1.5-flash', 'Gemini 1.5 Flash', 'chat', 1048576, 8192, 0.000075, 0.0003, '["text", "vision", "code"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'cohere'), 'command-r-plus', 'Command R+', 'chat', 128000, 4096, 0.003, 0.015, '["text", "code", "reasoning"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'cohere'), 'command-r', 'Command R', 'chat', 128000, 4096, 0.0005, 0.0015, '["text", "code"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'mistral'), 'mistral-large-latest', 'Mistral Large', 'chat', 128000, 4096, 0.002, 0.006, '["text", "code", "reasoning"]'::jsonb),
    ((SELECT id FROM providers WHERE name = 'mistral'), 'mistral-small-latest', 'Mistral Small', 'chat', 32000, 4096, 0.001, 0.003, '["text", "code"]'::jsonb)
ON CONFLICT (provider_id, name) DO NOTHING;