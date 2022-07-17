-- schema for amenities table
CREATE TYPE amenities_type AS ENUM(
    'power',
    'water',
    'gas',
    'parking',
    'air_condition',
    'internet',
    'heat',
    'tv',
    'swimming_pool',
    'garbage',
    'security',
    'lease_paperwork',
    'playground'
);
CREATE TABLE amenities(
    a_amenities_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    a_property_id UUID NOT NULL REFERENCES properties,
    a_amenity_type amenities_type NOT NULL DEFAULT NULL,
    a_currency TEXT NOT NULL DEFAULT NULL,
    a_amenity_fee TEXT NOT NULL DEFAULT NULL,
    a_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    a_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    a_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
    -- PRIMARY KEY(amenities_id, amenities, amenity_type)
);
-- create index for amenities
CREATE INDEX index_amenities_table
    ON amenities(a_amenities_id, a_property_id, a_amenity_type, a_created_at)
