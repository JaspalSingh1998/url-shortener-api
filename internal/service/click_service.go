package service

import (
	"context"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

type ClickService struct {
	store *store.ClickStore
}

func NewClickService(store *store.ClickStore) *ClickService {
	return &ClickService{store: store}
}

func (s *ClickService) Track(ctx context.Context, e *model.ClickEvent) {
	// fire and foreget
	_ = s.store.Create(ctx, e)
}
