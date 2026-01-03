package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisLinkCache struct {
	client *redis.Client
}

func NewRedisLinkCache(client *redis.Client) *RedisLinkCache {
	return &RedisLinkCache{
		client: client,
	}
}

func (c *RedisLinkCache) Get(ctx context.Context, shortCode string) (*model.Link, error) {
	key := fmt.Sprintf("link:%s", shortCode)

	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // cache miss
	}

	if err != nil {
		return nil, err
	}

	var link model.Link

	if err := json.Unmarshal([]byte(val), &link); err != nil {
		return nil, err
	}

	return &link, nil
}

func (c *RedisLinkCache) Set(ctx context.Context, link *model.Link, ttl time.Duration) error {
	key := fmt.Sprintf("link:%s", link.ShortCode)

	b, err := json.Marshal(link)

	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, b, ttl).Err()
}
