package store

import (
	"context"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkStore struct {
	db *pgxpool.Pool
}

func NewLinkStore(db *pgxpool.Pool) *LinkStore {
	return &LinkStore{
		db: db,
	}
}

func (s *LinkStore) Create(ctx context.Context, link *model.Link) error {
	query := `
		INSERT INTO links (short_code, original_url, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, is_active, created_at, updated_at
	`

	return s.db.QueryRow(
		ctx, query, link.ShortCode, link.OriginalURL, link.ExpiresAt,
	).Scan(
		&link.ID,
		&link.IsActive,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
}

func (s *LinkStore) GetByShortCode(ctx context.Context, shortCode string) (*model.Link, error) {
	query := `SELECT id, short_code, original_url, expires_at, is_active 
	FROM links
	WHERE short_code = $1
	AND is_active = true
	AND (expires_at IS NULL OR expires_at > NOW())
	`

	var link model.Link

	err := s.db.QueryRow(ctx, query, shortCode).Scan(
		&link.ID,
		&link.ShortCode,
		&link.OriginalURL,
		&link.ExpiresAt,
		&link.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return &link, nil
}
