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

func (s *AnalyticsStore) GetDailyStats(
	ctx context.Context,
	linkID int64,
	orgID int64,
	from, to time.Time,
) ([]DailyStat, error) {

	query := `
		SELECT d.date, d.clicks
		FROM link_click_stats_daily d
		WHERE d.link_id = $1
		  AND d.date >= $2
		  AND d.date <= $3
		  AND EXISTS (
			  SELECT 1
			  FROM links l
			  WHERE l.id = d.link_id
			    AND l.org_id = $4
		  )
		ORDER BY d.date
	`

	rows, err := s.db.Query(ctx, query, linkID, from, to, orgID)
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
	orgID int64,
	from, to time.Time,
) ([]HourlyStat, error) {

	query := `
		SELECT h.hour, h.clicks
		FROM link_click_stats_hourly h
		WHERE h.link_id = $1
		  AND h.hour >= $2
		  AND h.hour <= $3
		  AND EXISTS (
			  SELECT 1
			  FROM links l
			  WHERE l.id = h.link_id
			    AND l.org_id = $4
		  )
		ORDER BY h.hour
	`

	rows, err := s.db.Query(ctx, query, linkID, from, to, orgID)
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
