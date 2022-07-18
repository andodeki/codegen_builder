-- schema for users table
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    u_user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    u_email TEXT UNIQUE NOT NULL,
    u_phone TEXT NOT NULL DEFAULT NULL,
    u_password_hash bytea NOT NULL,
    u_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    u_updated_at TIMESTAMPTZ DEFAULT NULL,
    u_deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE UNIQUE INDEX index_user_table
    ON users (u_email);