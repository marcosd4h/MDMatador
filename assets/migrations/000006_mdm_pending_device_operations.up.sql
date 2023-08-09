CREATE TABLE mdm_pending_device_operations (
    device_id TEXT NOT NULL,
    cmd_verb TEXT NOT NULL,    
    setting_uri TEXT NOT NULL,
    setting_value TEXT NOT NULL,
    UNIQUE (device_id, cmd_verb, setting_uri, setting_value)
);