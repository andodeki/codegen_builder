CREATE TABLE contrators(
    c_contractor_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    c_business_id UUID NOT NULL REFERENCES businesses,
    c_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    c_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    c_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for contrators
CREATE INDEX index_contrators_table
    ON contrators(c_contractor_id, c_business_id, c_created_at)