CREATE TABLE properties(
    p_property_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    p_user_id UUID NOT NULL REFERENCES users,
    p_business_id UUID NOT NULL REFERENCES businesses,
    p_tenant_id UUID NOT NULL REFERENCES users,
    p_property_name TEXT NOT NULL DEFAULT NULL,
    p_property_type TEXT NOT NULL DEFAULT NULL,
    p_property_description TEXT NOT NULL DEFAULT NULL,
    p_property_doc TIMESTAMPTZ DEFAULT NULL, --business date of construction
    p_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), --data of listing
    p_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    p_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);
-- create index for properties
CREATE INDEX index_properties_table
    ON properties(p_property_id, p_property_type)
