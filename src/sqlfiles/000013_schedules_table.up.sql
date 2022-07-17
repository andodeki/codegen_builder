CREATE TYPE schedule_type AS ENUM(
    'connection',
    'disconnection',
    'interruption'
);
CREATE TABLE schedules(
    s_schedule_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    s_zone_id UUID NOT NULL REFERENCES zones,
    s_contractor_id UUID NOT NULL REFERENCES contrators,
    s_schedule_type schedule_type NOT NULL DEFAULT NULL,
    s_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    s_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    s_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for schedules
CREATE INDEX index_schedules_table
    ON schedules(s_schedule_id, s_zone_id, s_contractor_id, s_schedule_type, s_created_at)