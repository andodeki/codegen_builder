CREATE TYPE device_id AS ENUM(
    'device',
);
CREATE TYPE os_name AS ENUM(
    'os_name',
);
CREATE TYPE os_version AS ENUM(
    'os_version',
);
CREATE TYPE browser_name AS ENUM(
    'browser_name',
);
CREATE TYPE browser_version AS ENUM(
    'browser_version',
);
CREATE TYPE ip_type AS ENUM(
    'ipv4',
    'ipv6',
);
-- schema for sessions table
CREATE TABLE sessions(
    s_user_id UUID REFERENCES users,
    s_device_id device_id NOT NULL,
    s_os_name os_name NOT NULL,
    s_os_version os_version NOT NULL,
    s_browser_name browser_name NOT NULL,
    s_browser_version browser_version NOT NULL,
    s_ip ip_type NOT NULL,
    s_refresh_token TEXT NOT NULL,
    s_expires_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (s_user_id, s_device_id)
);