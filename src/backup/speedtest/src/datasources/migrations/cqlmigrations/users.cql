-- AUTO GENERATED By codegen_builder, DO NOT EDIT
--users
CREATE TABLE IF NOT EXISTS propertylisting.users( 
    u_user_id UUID, 
    u_email TEXT, 
    u_phone TEXT, 
    u_password_hash TEXT, 
    u_created_at TIMESTAMP, 
    u_updated_at TIMESTAMP, 
    u_deleted_at TIMESTAMP,
    PRIMARY KEY(u_user_id)
);

--u_user_id 
CREATE MATERIALIZED VIEW IF NOT EXISTS users_by_u_user_id AS 
  SELECT * FROM users
  WHERE u_user_id IS NOT NULL
  AND u_email IS NOT NULL
  AND u_phone IS NOT NULL
  AND u_password_hash IS NOT NULL
  AND u_created_at IS NOT NULL
  AND u_updated_at IS NOT NULL
  AND u_deleted_at IS NOT NULL
  PRIMARY KEY(u_user_id,u_email,u_phone,u_password_hash,u_created_at,u_updated_at,u_deleted_at);



CREATE INDEX index_user_table ON users(
    u_email
);
