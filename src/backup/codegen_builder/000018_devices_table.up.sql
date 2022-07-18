CREATE TABLE devices(
    d_device_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    d_user_id UUID NOT NULL REFERENCES users,
    
    d_device_name TEXT NOT NULL DEFAULT NULL,
    d_device_number TEXT NOT NULL DEFAULT NULL, -- meter number / surge protector

    d_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    d_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    d_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
);

-- create index for devices
CREATE INDEX index_devices_table
    ON devices(d_device_id, d_user_id, d_device_name, d_device_number, d_created_at)

-- /*

-- Device Table
-- ---d_device_id
-- ---d_user_id
-- ---d_device_name
-- ---d_status
-- Tanesco Admin

-- Payments Table
-- ---p_payment_id
-- ---p_token_id
-- ---p_user_id
-- ---p_business_id
-- ---p_amount
-- ---p_payment_type

-- =========================
-- TokensAccount Table
-- ---ta_token_id
-- ---ta_conn_id
-- ---ta_user_id
-- ---ta_business_id
-- ---ta_property_id
-- ---ta_token_amt
-- ---ta_token_bal

-- TokensEntries Table
-- ---te_id
-- ---te_token_acct_id
-- ---te_amt
-- ---te_units

-- TokenTransfers Table
-- ---tt_id
-- ---tt_from_token_acct
-- ---tt_to_token_acct
-- ---tt_token_units
-- =============================
-- Schedule Table
-- ---s_id
-- ---s_conn_id //fk
-- ---s_schedule_type // connectionSchedule, disconnectioShedule, interruptionShedule

-- Connection Table
-- ---c_conn_id
-- ---c_payment_id //fk
-- ---c_schedule_id //fk
-- ---c_zone_id
-- ---c_conn_type


-- Zone Table
-- ---z_zone_id
-- ---z_name
-- ---z_status  --up and running or down or meintenance
-- ---z_contractor_id

-- Contrators Table
-- ---c_contractor_id
-- ---c_business_id


-- Payments Table
-- ---p_payment_id
-- ---p_token_id
-- ---p_user_id
-- ---p_business_id
-- ---p_amount
-- ---p_payment_type




-- */