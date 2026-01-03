CREATE TABLE click_events (
    id BIGSERIAL PRIMARY KEY,

    link_id BIGINT NOT NULL,
    short_code VARCHAR(16) NOT NULL,

    ip_address INET NULL,
    user_agent TEXT NULL,
    referrer TEXT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_click_events_link_id
ON click_events (link_id);

CREATE INDEX idx_click_events_created_at
ON click_events (created_at);
