CREATE TABLE token_transfers(
    tt_transfer_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tt_from_conn_acct UUID NOT NULL REFERENCES connection_acct,
    tt_to_conn_acct UUID NOT NULL REFERENCES connection_acct,

    tt_token_units BIGINT NOT NULL,
    tt_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    tt_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    tt_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for token_transfers
CREATE INDEX index_token_transfers_table
    ON token_transfers(tt_transfer_id, tt_from_conn_acct, tt_to_conn_acct, tt_token_units, tt_created_at)