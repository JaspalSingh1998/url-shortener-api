package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DailyStat struct {
	Date   time.Time
	Clicks int64
}

type HourlyStat struct {
	Hour   time.Time
	Clicks int64
}

type AnalyticsStore struct {
	db *pgxpool.Pool
}

func NewAnalyticsStore(db *pgxpool.Pool) *AnalyticsStore {
	return &AnalyticsStore{db: db}
}

func (s *AnalyticsStore) GetDailyStats(ctx context.Context, linkID int64, from, to time.Time) ([]DailyStat, error) {
	query := `
		SELECT date, clicks
		FROM link_click_stats_daily
		WHERE link_id = $1
		  AND date >= $2
		  AND date <= $3
		ORDER BY date
	`
	rows, err := s.db.Query(ctx, query, linkID, from, to)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []DailyStat
	for rows.Next() {
		var s DailyStat
		if err := rows.Scan(&s.Date, &s.Clicks); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return stats, nil
}

func (s *AnalyticsStore) GetHourlyStats(
	ctx context.Context,
	linkID int64,
	from, to time.Time,
) ([]HourlyStat, error) {

	query := `
		SELECT hour, clicks
		FROM link_click_stats_hourly
		WHERE link_id = $1
		  AND hour >= $2
		  AND hour <= $3
		ORDER BY hour
	`

	rows, err := s.db.Query(ctx, query, linkID, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []HourlyStat
	for rows.Next() {
		var s HourlyStat
		if err := rows.Scan(&s.Hour, &s.Clicks); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return stats, nil
}
