-- Create billing_transactions table
CREATE TABLE IF NOT EXISTS billing_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
    billing_account_id UUID REFERENCES billing_accounts(id) ON DELETE CASCADE,
    api_request_id UUID REFERENCES api_requests(id) ON DELETE SET NULL,
    transaction_type VARCHAR(50) NOT NULL,
    amount DECIMAL(10,6) NOT NULL,
    currency VARCHAR(3) DEFAULT 'USD',
    description TEXT,
    metadata JSONB DEFAULT '{}'::jsonb,
    balance_after DECIMAL(10,2) NOT NULL,
);

-- Create indexes for billing_transactions
CREATE INDEX IF NOT EXISTS idx_billing_transactions_billing_account_id ON billing_transactions(billing_account_id);
CREATE INDEX IF NOT EXISTS idx_billing_transactions_api_request_id ON billing_transactions(api_request_id);
CREATE INDEX IF NOT EXISTS idx_billing_transactions_transaction_type ON billing_transactions(transaction_type);
CREATE INDEX IF NOT EXISTS idx_billing_transactions_created_at ON billing_transactions(created_at);
CREATE INDEX IF NOT EXISTS idx_billing_transactions_amount ON billing_transactions(amount);