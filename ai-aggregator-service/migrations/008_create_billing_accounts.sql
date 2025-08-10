-- Create billing_accounts table
CREATE TABLE IF NOT EXISTS billing_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    organization_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    account_type VARCHAR(50) NOT NULL DEFAULT 'organization',
    balance DECIMAL(10,2) DEFAULT 0.0,
    currency VARCHAR(3) DEFAULT 'USD',
    billing_email VARCHAR(255),
    stripe_customer_id VARCHAR(255),
    stripe_subscription_id VARCHAR(255),
    subscription_status VARCHAR(50) DEFAULT 'active',
    subscription_tier VARCHAR(50) DEFAULT 'free',
    monthly_budget DECIMAL(10,2),
    alert_threshold DECIMAL(10,2) DEFAULT 10.0,
    is_active BOOLEAN DEFAULT TRUE,
    CONSTRAINT check_owner_type CHECK (
        (account_type = 'organization' AND organization_id IS NOT NULL AND user_id IS NULL) OR
        (account_type = 'user' AND user_id IS NOT NULL AND organization_id IS NULL)
    )
);

-- Create indexes for billing_accounts
CREATE INDEX IF NOT EXISTS idx_billing_accounts_organization_id ON billing_accounts(organization_id);
CREATE INDEX IF NOT EXISTS idx_billing_accounts_user_id ON billing_accounts(user_id);
CREATE INDEX IF NOT EXISTS idx_billing_accounts_stripe_customer_id ON billing_accounts(stripe_customer_id);
CREATE INDEX IF NOT EXISTS idx_billing_accounts_subscription_status ON billing_accounts(subscription_status);
CREATE INDEX IF NOT EXISTS idx_billing_accounts_is_active ON billing_accounts(is_active);
CREATE INDEX IF NOT EXISTS idx_billing_accounts_created_at ON billing_accounts(created_at);