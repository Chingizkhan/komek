package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	cache "komek/pkg/redis_cache"
	"komek/pkg/redis_cache/rediscache"
	"sync"
)

type Cache struct {
	m      *sync.RWMutex
	c      *rediscache.RedisCache
	prefix string
}

func New(c *rediscache.RedisCache) *Cache {
	return &Cache{
		m:      &sync.RWMutex{},
		c:      c,
		prefix: "fundraise_",
	}
}

func (c *Cache) SetDonorsQuantityByFundraiseID(ctx context.Context, fundraiseID uuid.UUID) (int, error) {
	c.m.Lock()
	defer c.m.Unlock()

	var res int
	err := c.c.Get(ctx, c.getDonorsQuantityKey(fundraiseID), &res)
	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		return 0, fmt.Errorf("c.c.Get: %w", err)
	}

	err = c.c.Set(ctx, c.getDonorsQuantityKey(fundraiseID), res+1)
	if err != nil {
		return 0, fmt.Errorf("c.c.Get: %w", err)
	}

	return res + 1, nil
}

func (c *Cache) GetDonorsQuantityByFundraiseID(ctx context.Context, fundraiseID uuid.UUID) (int, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	var res int
	err := c.c.Get(ctx, c.getDonorsQuantityKey(fundraiseID), &res)
	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		return 0, fmt.Errorf("c.c.Get: %w", err)
	}

	return res, nil
}

func (c *Cache) getDonorsQuantityKey(fundraiseID uuid.UUID) string {
	return c.prefix + "donors_quantity_" + fundraiseID.String()
}
