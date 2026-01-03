package service

import (
	"context"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

type AggregationService struct {
	store *store.AggregationStore
}

func NewAggregationService(store *store.AggregationStore) *AggregationService {
	return &AggregationService{store: store}
}

func (s *AggregationService) RunHourly(ctx context.Context, t time.Time) error {
	from := t.Truncate(time.Hour)
	to := from.Add(time.Hour)

	return s.store.AggregateHourly(ctx, from, to)
}

func (s *AggregationService) RunDaily(ctx context.Context, t time.Time) error {
	from := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	to := from.Add(24 * time.Hour)
	return s.store.AggregateDaily(ctx, from, to)
}
