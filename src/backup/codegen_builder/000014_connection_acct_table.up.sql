CREATE TYPE connection_type AS ENUM(
    'prepaid',
    'postpaid'
);

CREATE TABLE connection_acct( -- connection account
    ta_conn_acct_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ta_user_id UUID NOT NULL REFERENCES users,
    ta_business_id UUID NOT NULL REFERENCES businesses,
    ta_property_id UUID NOT NULL REFERENCES properties,
    ta_zone_id UUID NOT NULL REFERENCES zones,
    ta_schedule_id UUID NOT NULL REFERENCES schedules,

    ta_conn_type connection_type NOT NULL DEFAULT NULL,
    ta_conn_amt BIGINT NOT NULL,
    ta_token_amt BIGINT NOT NULL,
    ta_token_bal BIGINT NOT NULL,
    ta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), --data of listing
    ta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    ta_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for connection_acct
CREATE INDEX index_connection_acct_table
    ON connection_acct(ta_conn_acct_id, ta_user_id, ta_business_id, ta_property_id, ta_conn_type, ta_created_at)