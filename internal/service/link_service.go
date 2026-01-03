package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/cache"
	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

var ErrLinkNotFound = errors.New("link not found or expired")

type LinkService struct {
	store *store.LinkStore
	cache cache.LinkCache
}

func NewLinkService(store *store.LinkStore, cache cache.LinkCache) *LinkService {
	return &LinkService{
		store: store,
		cache: cache,
	}
}

func (s *LinkService) CreateLink(ctx context.Context, originalUrl string, customAlias string, expiresAt *time.Time) (*model.Link, error) {
	if expiresAt != nil && expiresAt.Before(time.Now()) {
		return nil, errors.New("expires_at must be in the future")
	}

	shortCode := customAlias
	if shortCode == "" {
		shortCode = generateShortCode(7)
	}

	link := &model.Link{
		ShortCode:   shortCode,
		OriginalURL: originalUrl,
		ExpiresAt:   expiresAt,
	}

	if err := s.store.Create(ctx, link); err != nil {
		return nil, err
	}

	return link, nil
}

func (s *LinkService) ResolveLink(ctx context.Context, shortCode string) (*model.Link, error) {
	// 1 Cache Lookup
	if s.cache != nil {
		if link, _ := s.cache.Get(ctx, shortCode); link != nil {
			return link, nil
		}
	}

	// 2 DB fallback
	link, err := s.store.GetByShortCode(ctx, shortCode)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLinkNotFound
		}
		return nil, err
	}

	// 3 Cache Population
	if s.cache != nil {
		ttl := 24 * time.Hour
		if link.ExpiresAt != nil {
			ttl = time.Until(*link.ExpiresAt)
		}
		_ = s.cache.Set(ctx, link, ttl)
	}

	return link, nil
}
