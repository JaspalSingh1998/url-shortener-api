CREATE TABLE links (
    id BIGSERIAL PRIMARY KEY,

    short_code VARCHAR(16) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT true,
    expires_at TIMESTAMPTZ NULL,

    created_by BIGINT NULL,
    org_id BIGINT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_links_short_code
ON links (short_code);

CREATE INDEX idx_links_active_expiry
ON links (is_active, expires_at);
