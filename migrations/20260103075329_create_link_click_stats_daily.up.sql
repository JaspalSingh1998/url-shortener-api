CREATE TABLE link_click_stats_daily (
    link_id BIGINT NOT NULL,
    short_code VARCHAR(16) NOT NULL,
    date DATE NOT NULL,

    clicks BIGINT NOT NULL DEFAULT 0,

    PRIMARY KEY (link_id, date)
);

CREATE INDEX idx_link_click_stats_daily_date
ON link_click_stats_daily (date);
