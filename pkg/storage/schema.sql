CREATE TABLE IF NOT EXISTS tenents (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS auth (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    api_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenents(id) ON DELETE CASCADE,
    CONSTRAINT auth_tenant_api_key_unique UNIQUE (tenant_id, api_key)
);

CREATE TABLE IF NOT EXISTS secrets (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    secret_key VARCHAR(255) NOT NULL,
    secret_value BYTEA NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dek_version INTEGER NOT NULL,
    version INTEGER NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenents(id) ON DELETE CASCADE,
    CONSTRAINT secrets_tenant_key_version_unique UNIQUE (tenant_id, secret_key, version)
);

CREATE TABLE IF NOT EXISTS deks (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    dek BYTEA NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    version INTEGER NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenents(id) ON DELETE CASCADE,
    CONSTRAINT deks_tenant_version_unique UNIQUE (tenant_id, version)
);

CREATE INDEX IF NOT EXISTS idx_auth_tenant_id ON auth(tenant_id, api_key);
CREATE INDEX IF NOT EXISTS idx_secrets_tenant_id ON secrets(tenant_id, secret_key, version);
CREATE INDEX IF NOT EXISTS idx_deks_tenant_id ON deks(tenant_id, version);