-- schema for sessions table
CREATE TABLE sessions(
    s_user_id UUID REFERENCES users,
    s_device_id TEXT NOT NULL,
    s_os_name TEXT NOT NULL,
    s_os_version TEXT NOT NULL,
    s_browser_name TEXT NOT NULL,
    s_browser_version TEXT NOT NULL,
    s_ip TEXT NOT NULL,
    s_refresh_token TEXT NOT NULL,
    s_expires_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (s_user_id, s_device_id)
);