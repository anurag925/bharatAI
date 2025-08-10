-- Create api_requests table
CREATE TABLE IF NOT EXISTS api_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    api_key_id UUID REFERENCES api_keys(id) ON DELETE SET NULL,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    organization_id UUID REFERENCES organizations(id) ON DELETE SET NULL,
    provider_id UUID REFERENCES providers(id) ON DELETE SET NULL,
    model_id UUID REFERENCES models(id) ON DELETE SET NULL,
    request_id VARCHAR(255) UNIQUE NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    method VARCHAR(10) NOT NULL,
    endpoint VARCHAR(500) NOT NULL,
    headers JSONB DEFAULT '{}'::jsonb,
    request_body JSONB,
    response_body JSONB,
    error_message TEXT,
    status_code INTEGER,
    input_tokens INTEGER DEFAULT 0,
    output_tokens INTEGER DEFAULT 0,
    total_tokens INTEGER DEFAULT 0,
    cost DECIMAL(10,6) DEFAULT 0.0,
    latency_ms INTEGER,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes for api_requests
CREATE INDEX IF NOT EXISTS idx_api_requests_api_key_id ON api_requests(api_key_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_user_id ON api_requests(user_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_organization_id ON api_requests(organization_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_provider_id ON api_requests(provider_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_model_id ON api_requests(model_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_request_id ON api_requests(request_id);
CREATE INDEX IF NOT EXISTS idx_api_requests_status ON api_requests(status);
CREATE INDEX IF NOT EXISTS idx_api_requests_created_at ON api_requests(created_at);
CREATE INDEX IF NOT EXISTS idx_api_requests_completed_at ON api_requests(completed_at);
CREATE INDEX IF NOT EXISTS idx_api_requests_ip_address ON api_requests(ip_address);