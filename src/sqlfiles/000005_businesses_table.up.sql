CREATE TABLE businesses(
    b_business_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    b_user_id UUID NOT NULL REFERENCES users,
    b_department TEXT NOT NULL DEFAULT NULL,
    b_business_name TEXT NOT NULL DEFAULT NULL,
    b_business_type TEXT NOT NULL DEFAULT NULL,
    b_business_description TEXT NOT NULL DEFAULT NULL,
    b_business_dof TIMESTAMPTZ DEFAULT NULL, --business date of formation
    b_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    b_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    b_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
    -- PRIMARY KEY(business_id, businesses)
);

-- create index for businesses
CREATE INDEX index_businesses_table
    ON businesses(b_business_id, b_business_type)