package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

var ErrLinkNotFound = errors.New("link not found or expired")

type LinkService struct {
	store *store.LinkStore
}

func NewLinkService(store *store.LinkStore) *LinkService {
	return &LinkService{
		store: store,
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
	link, err := s.store.GetByShortCode(ctx, shortCode)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLinkNotFound
		}
		return nil, err
	}

	return link, nil
}
