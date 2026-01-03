CREATE TABLE link_click_stats_hourly (
    link_id BIGINT NOT NULL,
    short_code VARCHAR(16) NOT NULL,
    hour TIMESTAMPTZ NOT NULL,

    clicks BIGINT NOT NULL DEFAULT 0,

    PRIMARY KEY (link_id, hour)
);

CREATE INDEX idx_link_click_stats_hourly_hour
ON link_click_stats_hourly (hour);
