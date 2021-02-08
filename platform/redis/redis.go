package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/royalbhati/urlshortener/cmd/config"
)

func Open(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Address + ":" + cfg.Redis.Port,
		DB:   cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil

}
