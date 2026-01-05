package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	redis *redis.Client
}

func NewRateLimiter(redis *redis.Client) *RateLimiter {
	return &RateLimiter{redis: redis}
}

func (r *RateLimiter) Allow(
	ctx context.Context,
	key string,
	limit int,
	window time.Duration,
) (bool, error) {
	pipe := r.redis.TxPipeline()

	count := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, window)

	_, err := pipe.Exec(ctx)

	if err != nil {
		return false, err
	}

	return count.Val() <= int64(limit), nil
}

func (r *RateLimiter) LimitByIP(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rl:ip:%s", ip)

		allowed, err := r.Allow(c.Request.Context(), key, limit, window)

		if err != nil || !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exeeded",
			})
			return
		}

		c.Next()
	}
}

func (r *RateLimiter) LimitByOrg(
	limit int, window time.Duration,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*Claims)
		orgID := claims.OrgID

		key := fmt.Sprintf("rl:org:%s", orgID)

		allowed, err := r.Allow(c.Request.Context(), key, limit, window)
		if err != nil || !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exeeded",
			})

			return
		}
		c.Next()
	}
}
