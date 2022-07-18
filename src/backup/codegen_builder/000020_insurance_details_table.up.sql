CREATE TABLE insurance_details(
    idt_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    idt_user_id UUID NOT NULL REFERENCES users,
    idt_asset_id UUID NOT NULL REFERENCES assets,
    idt_insurance_type TEXT NOT NULL,
    idt_basic_cover JSON NOT NULL,
    idt_add_ons JSON NOT NULL,
    idt_start_date TIMESTAMPTZ NOT NULL,
    idt_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    idt_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    idt_deleted_at TIMESTAMPTZ
);

-- create index for insurance_details
CREATE INDEX index_insurance_details_table
    ON insurance_details(idt_id, idt_user_id, idt_insurance_type, idt_start_date)
