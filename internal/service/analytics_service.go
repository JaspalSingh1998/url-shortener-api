package service

import (
	"context"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

type AnalyticsService struct {
	store *store.AnalyticsStore
}

func NewAnalyticsService(store *store.AnalyticsStore) *AnalyticsService {
	return &AnalyticsService{store: store}
}

func (s *AnalyticsService) Daily(
	ctx context.Context,
	linkID int64,
	from, to time.Time,
) ([]store.DailyStat, error) {
	return s.store.GetDailyStats(ctx, linkID, from, to)
}

func (s *AnalyticsService) Hourly(
	ctx context.Context,
	linkID int64,
	from, to time.Time,
) ([]store.HourlyStat, error) {
	return s.store.GetHourlyStats(ctx, linkID, from, to)
}
