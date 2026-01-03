package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AggregationStore struct {
	db *pgxpool.Pool
}

func NewAggregationStore(db *pgxpool.Pool) *AggregationStore {
	return &AggregationStore{db: db}
}

func (s *AggregationStore) AggregateDaily(ctx context.Context, from, to time.Time) error {
	query := `
	INSERT INTO link_click_stats_daily (link_id, short_code, date, clicks)
		SELECT
			link_id,
			short_code,
			DATE(created_at) AS date,
			COUNT(*) AS clicks
		FROM click_events
		WHERE created_at >= $1
		AND created_at < $2
		GROUP BY link_id, short_code, DATE(created_at)
		ON CONFLICT (link_id, date)
		DO UPDATE SET clicks = link_click_stats_daily.clicks + EXCLUDED.clicks;
	`

	_, err := s.db.Exec(ctx, query, from, to)

	return err
}

func (s *AggregationStore) AggregateHourly(
	ctx context.Context,
	from, to time.Time,
) error {
	query := `
		INSERT INTO link_click_stats_hourly (link_id, short_code, hour, clicks)
		SELECT
			link_id,
			short_code,
			date_trunc('hour', created_at) AS hour,
			COUNT(*) AS clicks
		FROM click_events
		WHERE created_at >= $1
		AND created_at < $2
		GROUP BY link_id, short_code, date_trunc('hour', created_at)
		ON CONFLICT (link_id, hour)
		DO UPDATE SET clicks = link_click_stats_hourly.clicks + EXCLUDED.clicks;
	`

	_, err := s.db.Exec(ctx, query, from, to)

	return err
}
