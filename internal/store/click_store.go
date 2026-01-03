package store

import (
	"context"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClickStore struct {
	db *pgxpool.Pool
}

func NewClickStore(db *pgxpool.Pool) *ClickStore {
	return &ClickStore{
		db: db,
	}
}

func (s *ClickStore) Create(ctx context.Context, e *model.ClickEvent) error {
	query := `
		INSERT INTO click_events (
		link_id, short_code, ip_address, user_agent, referrer
		) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.db.Exec(
		ctx,
		query,
		e.LinkID,
		e.ShortCode,
		e.IPAddress,
		e.UserAgent,
		e.Referrer,
	)

	return err
}
