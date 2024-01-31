package locker

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Locker struct {
	c           *redis.Client
	lockTimeout time.Duration
}

func New(c *redis.Client, lockTimeout time.Duration) *Locker {
	return &Locker{c, lockTimeout}
}

func (s Locker) Lock(key string) error {
	isSet, err := s.c.Exists(key).Result()
	if err != nil {
		return fmt.Errorf("s.c.Get: %w", err)
	}
	exists := isSet != 0
	if exists {
		return errors.New(fmt.Sprintf("key 'lock-%s' is locked by other process", key))
	}

	_, err = s.c.Set("lock-"+key, "val", s.lockTimeout).Result()
	if err != nil {
		return fmt.Errorf("exists s.c.Set: %w", err)
	}

	return nil
}

func (s Locker) Unlock(key string) error {
	_, err := s.c.Del("lock-" + key).Result()
	if err != nil {
		return fmt.Errorf("s.c.Del: %w", err)
	}
	return nil
}
