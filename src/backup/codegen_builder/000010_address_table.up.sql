CREATE TABLE address(
    a_address_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    a_user_id UUID NOT NULL REFERENCES users,
    a_business_id UUID NOT NULL REFERENCES businesses,
    a_property_id UUID NOT NULL REFERENCES properties,
    a_lat TEXT NOT NULL DEFAULT NULL,
    a_long TEXT NOT NULL DEFAULT NULL,
    a_city TEXT NOT NULL DEFAULT NULL,
    a_province TEXT NOT NULL DEFAULT NULL,
    a_country TEXT NOT NULL DEFAULT NULL,
    a_plotno TEXT NOT NULL DEFAULT NULL,
    a_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    a_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    a_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
    -- PRIMARY KEY(address_id, address)
);

-- create index for address
CREATE INDEX index_address_table
    ON address(a_address_id, a_user_id, a_business_id, a_property_id, a_lat, a_long)