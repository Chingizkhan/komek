package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	connAttempts int
	connTimeout  time.Duration
	Client       *redis.Client
}

func New(url, password string, opts ...Option) (*Redis, error) {
	r := &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
		}),
	}

	_, err := r.Client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("r.Client.Ping(): %w", err)
	}

	for _, opt := range opts {
		opt(r)
	}

	return r, nil
}
