-- schema for units table
CREATE TYPE owner_type AS ENUM(
    'rental',
    'buying',
    'leasing'
);
CREATE TABLE units(
    u_unit_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    u_property_id UUID NOT NULL REFERENCES properties,
    u_tenant_id UUID NOT NULL REFERENCES users,
    u_space_units_unit TEXT NOT NULL DEFAULT NULL, --units_
    u_space_units_type TEXT NOT NULL DEFAULT NULL,
    u_space_units_capacity TEXT NOT NULL DEFAULT NULL,
    u_space_units_no TEXT NOT NULL DEFAULT NULL,
    u_space_units_flr_level TEXT NOT NULL DEFAULT NULL,
    u_space_units_sq_area TEXT NOT NULL DEFAULT NULL,
    u_space_units_flr_plans TEXT NOT NULL DEFAULT NULL,
    u_furnished BOOLEAN NOT NULL DEFAULT NULL,
    u_refurbishing BOOLEAN NOT NULL DEFAULT NULL,
    u_bronchure_uploads TEXT NOT NULL DEFAULT NULL,
    u_ownership_type owner_type NOT NULL DEFAULT NULL,
    u_ownership_docs TEXT NOT NULL DEFAULT NULL,
    u_occupancy_docs TEXT NOT NULL DEFAULT NULL,
    u_currency TEXT NOT NULL DEFAULT NULL,
    u_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), --date of listing
    u_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    u_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
    -- PRIMARY KEY(property_id, properties, property_type)
);
-- create index for units
CREATE INDEX index_units_table
    ON units(u_unit_id, u_space_units_unit, u_furnished, u_space_units_capacity)
