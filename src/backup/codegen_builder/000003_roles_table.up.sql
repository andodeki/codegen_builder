CREATE TYPE user_role AS ENUM(
    'admin'
    -- 'member',
    -- 'memberIsTarget',
    -- 'anonym',
);

-- schema for roles table
CREATE TABLE roles(
    r_user_id UUID NOT NULL REFERENCES users,
    r_role user_role NOT NULL,
    r_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(r_user_id, r_role)
);

-- create index for roles
CREATE INDEX index_roles_table
    ON roles(r_user_id)
