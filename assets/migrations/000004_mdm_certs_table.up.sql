CREATE TABLE mdm_certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    identity_cert BLOB NOT NULL,
    identity_key BLOB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP    
);

CREATE INDEX idx_mdm_certificates_id ON mdm_certificates(id);
