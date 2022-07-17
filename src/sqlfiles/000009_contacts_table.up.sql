CREATE TABLE contacts(
    c_contact_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    c_user_id UUID NOT NULL REFERENCES users,
    c_business_id UUID NOT NULL REFERENCES businesses,
    c_property_id UUID NOT NULL REFERENCES properties,
    c_website TEXT NOT NULL DEFAULT NULL,
    c_email TEXT NOT NULL DEFAULT NULL,
    c_email_verified BOOLEAN NOT NULL DEFAULT NULL,
    c_other_email TEXT NOT NULL DEFAULT NULL,
    c_mobile_phone TEXT NOT NULL DEFAULT NULL,
    c_phone_verified BOOLEAN NOT NULL DEFAULT NULL,
    c_other_phone TEXT NOT NULL DEFAULT NULL,
    c_facebook_links TEXT NOT NULL DEFAULT NULL,
    c_youtube_links TEXT NOT NULL DEFAULT NULL,
    c_twitter_links TEXT NOT NULL DEFAULT NULL,
    c_instagram_links TEXT NOT NULL DEFAULT NULL,
    c_linkedin_links TEXT NOT NULL DEFAULT NULL,
    c_skype_links TEXT NOT NULL DEFAULT NULL,
    c_snapchat_links TEXT NOT NULL DEFAULT NULL,
    c_telegram_links TEXT NOT NULL DEFAULT NULL,
    c_whatsapp_links TEXT NOT NULL DEFAULT NULL,
    c_wechat_links TEXT NOT NULL DEFAULT NULL,
    c_discord_links TEXT NOT NULL DEFAULT NULL,
    c_slack_links TEXT NOT NULL DEFAULT NULL,
    c_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    c_updated_at TIMESTAMPTZ NOT NULL DEFAULT NULL,
    c_deleted_at TIMESTAMPTZ NOT NULL DEFAULT NULL
    -- PRIMARY KEY(contact_id, contacts)
);

-- create index for contacts
CREATE INDEX index_contacts_table
    ON contacts(c_contact_id, c_user_id, c_email, c_mobile_phone, c_created_at)