package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisStore struct {
	Client *redis.Client
}

func (s *RedisStore) Set(ctx context.Context, key string, data []byte) error {
	status := s.Client.Set(ctx, key, data, time.Hour*24)
	return status.Err()
}

func (s *RedisStore) Get(ctx context.Context, key string) ([]byte, error) {
	result := s.Client.Get(ctx, key)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return []byte(result.Val()), nil
}

func NewRedisStore(options *redis.Options) (*RedisStore, error) {
	client := redis.NewClient(options)

	status := client.Ping(context.Background())
	if status.Err() != nil {
		return nil, status.Err()
	}

	return &RedisStore{
		Client: client,
	}, nil
}
