CREATE TYPE payments_channels AS ENUM(
    'mpesa_payments',
    't_Kash_payments',
    'airtel_Money_payments',
    'debit_credit_prepaid_payments',
    'eazzy_pay_payments',
    'e_agent_payments',
    'kcb_cash_payments',
    'equity_payments',
    'pesalink_payments',
    'paypal_payments',
    'cash_payments'
);
CREATE TYPE transaction_status AS ENUM(
    'new',
    'cancelled',
    'failed',
    'pending',
    'declined',
    'rejected',
    'success'
);
CREATE TYPE payments_type AS ENUM(
    'rental_fee',
    'buying_fee',
    'leasing_fee',
    'amenities_fee',
    'power_conn_fee',
    'water_conn_fee',
    'power_token_fee',
    'water_token_fee',
    'viewing_fee'
);
CREATE TYPE transaction_mode AS ENUM(
    'offline',
    'online',
    'wired',
    'draft',
    'cheque',
    'cash_on_delivery'
);
CREATE TYPE transaction_type AS ENUM(
    'credit',
    'debit'
);
CREATE TABLE payments(
    p_payment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    p_conn_acct_id UUID NOT NULL REFERENCES connection_acct,
    p_token_entry_id UUID NOT NULL REFERENCES token_entries,
    p_user_id UUID NOT NULL REFERENCES users,
    p_business_id UUID NOT NULL REFERENCES businesses,
    p_schedule_id UUID NOT NULL REFERENCES schedules,
    
    p_code TEXT NOT NULL,
    p_type transaction_type NOT NULL DEFAULT NULL,
    p_amount BIGINT NOT NULL,
    p_status transaction_status NOT NULL DEFAULT NULL,
    p_payment_type payments_type NOT NULL DEFAULT NULL,
    p_mode transaction_mode NOT NULL DEFAULT NULL,
    p_payment_channel payments_channels NOT NULL DEFAULT NULL,
    
    p_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    p_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    p_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for payments
CREATE INDEX index_payments_table
    ON payments(p_payment_id, p_conn_acct_id, p_user_id, p_business_id, p_amount, p_status, p_created_at)