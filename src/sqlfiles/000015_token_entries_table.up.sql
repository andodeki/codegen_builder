CREATE TABLE token_entries(
    te_token_entry_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    te_conn_acct_id UUID NOT NULL REFERENCES connection_acct,

    te_amt BIGINT NOT NULL,
    te_units BIGINT NOT NULL,
    te_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    te_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    te_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);
-- create index for token_entries
CREATE INDEX index_token_entries_table
    ON token_entries(te_token_entry_id, te_conn_acct_id, te_amt, te_units, te_created_at)