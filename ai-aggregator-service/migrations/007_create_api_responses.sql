-- Create api_responses table
CREATE TABLE IF NOT EXISTS api_responses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
    request_id UUID REFERENCES api_requests(id) ON DELETE CASCADE,
    provider_id UUID REFERENCES providers(id) ON DELETE SET NULL,
    model_id UUID REFERENCES models(id) ON DELETE SET NULL,
    response_data JSONB,
    usage_data JSONB,
    error_data JSONB,
    status_code INTEGER,
    latency_ms INTEGER,
    retry_count INTEGER DEFAULT 0,
);

-- Create indexes for api_responses
CREATE INDEX IF NOT EXISTS idx_api_responses_request_id ON api_responses(request_id);
CREATE INDEX IF NOT EXISTS idx_api_responses_provider_id ON api_responses(provider_id);
CREATE INDEX IF NOT EXISTS idx_api_responses_model_id ON api_responses(model_id);
CREATE INDEX IF NOT EXISTS idx_api_responses_status_code ON api_responses(status_code);
CREATE INDEX IF NOT EXISTS idx_api_responses_created_at ON api_responses(created_at);