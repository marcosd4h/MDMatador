CREATE TABLE mdm_enrollments (
    device_id   TEXT NOT NULL PRIMARY KEY,
    hardware_id TEXT NOT NULL,
    device_name TEXT NOT NULL,
    device_type TEXT NOT NULL,    
    os_locale   TEXT NOT NULL,
    os_edition  TEXT NOT NULL,
    os_version  TEXT NOT NULL,
    client_version TEXT NOT NULL,
    last_seen      TEXT NOT NULL,    
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL
);

CREATE INDEX idx_mdm_enrollments_device_id ON mdm_enrollments(device_id);
