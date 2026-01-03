package cache

import (
	"context"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
)

type LinkCache interface {
	Get(ctx context.Context, shortCode string) (*model.Link, error)
	Set(ctx context.Context, link *model.Link, ttl time.Duration) error
}
