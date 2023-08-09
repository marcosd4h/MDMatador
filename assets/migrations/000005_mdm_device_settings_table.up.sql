CREATE TABLE mdm_device_settings (
    device_id TEXT NOT NULL,
    setting_uri TEXT NOT NULL,
    setting_value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (device_id, setting_uri)
);
