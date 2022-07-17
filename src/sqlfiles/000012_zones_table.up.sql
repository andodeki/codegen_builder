CREATE TYPE status_type AS ENUM(
    'up',
    'down',
    'maintainance'
);
CREATE TABLE zones(
    z_zone_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    z_contractor_id UUID NOT NULL REFERENCES contrators,
    z_name TEXT NOT NULL DEFAULT NULL,
    z_status status_type NOT NULL DEFAULT NULL,
    z_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    z_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    z_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for zones
CREATE INDEX index_zones_table
    ON zones(z_zone_id, z_contractor_id, z_status)