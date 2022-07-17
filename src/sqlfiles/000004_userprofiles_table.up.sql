-- schema for userprofiles table
CREATE TABLE user_profiles(
    up_user_id UUID NOT NULL REFERENCES users,
    up_display_name TEXT NOT NULL DEFAULT NULL,
    up_first_name TEXT NOT NULL DEFAULT NULL,
    up_middle_name TEXT NOT NULL DEFAULT NULL,
    up_last_name TEXT NOT NULL DEFAULT NULL,
    up_DOB TIMESTAMPTZ NOT NULL DEFAULT NULL,
    up_gender BOOLEAN NOT NULL DEFAULT NULL,
    up_picture TEXT NOT NULL DEFAULT NULL,
    up_blocked BOOLEAN NOT NULL DEFAULT NULL,
    up_last_ip TEXT NOT NULL DEFAULT NULL,
    up_last_password_reset TIMESTAMPTZ NOT NULL DEFAULT NULL,
    up_last_login TIMESTAMPTZ NOT NULL DEFAULT NULL,
    up_payment_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    up_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    up_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    PRIMARY KEY(up_user_id, up_display_name)
);

-- create index for user_profiles
CREATE INDEX index_user_profiles_table
    ON user_profiles(up_user_id, up_display_name, up_first_name, up_middle_name)