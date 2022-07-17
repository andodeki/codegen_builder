CREATE TABLE assets(
    a_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    a_user_id UUID NOT NULL REFERENCES users,
    a_value TEXT NOT NULL,
    a_make TEXT NOT NULL,
    a_model TEXT NOT NULL,
    a_year_of_manufacture TIMESTAMPTZ NOT NULL,
    a_use TEXT NOT NULL,
    a_policy_period TIMESTAMPTZ NOT NULL,
    a_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    a_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    a_deleted_at TIMESTAMPTZ
);

-- create index for units
CREATE INDEX index_assets_table
    ON assets(a_id, a_model, a_year_of_manufacture, a_policy_period)
