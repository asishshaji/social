package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	SetAuthenticatedUser(context.Context, string) error
	PurgeAuthenticatedUser(context.Context, string) error
}

type cache struct {
	redis *redis.Client
}

func NewRedisCache(opts *redis.Options) ICache {
	return cache{
		redis: redis.NewClient(opts),
	}
}

func (c cache) SetAuthenticatedUser(ctx context.Context, username string) error {
	return c.redis.SAdd(ctx, "users:auth_issued", username).Err()
}

func (c cache) PurgeAuthenticatedUser(ctx context.Context, username string) error {
	return c.redis.SRem(ctx, "users:auth_issued", username).Err()
}
