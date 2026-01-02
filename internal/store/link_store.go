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
