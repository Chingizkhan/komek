package cache

import (
	"context"
	"errors"
	"time"
)

const (
	// NoTTL declares that cache does not have time to live, so cache do not expire.
	NoTTL = 0
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache interface {
	// Exists checks if cache key exists.
	Exists(ctx context.Context, key string) (bool, error)

	// Get gets cached data for a specific key.
	// Use pointers to get the value.
	Get(ctx context.Context, key string, value any) error

	// Set caches data for a specific key with default setting and TTL.
	Set(ctx context.Context, key string, value any) error

	// SetNX caches data for a specific key with default TTL.
	// NX -- Only set the key if it does not already exist.
	// returns true if set.
	SetNX(ctx context.Context, key string, value any) (bool, error)

	// Ensure gets cached data or sets if data does not exist.
	// value: is setting and getting value. Use pointers to get the value.
	// exist: is true when cache key existed before.
	Ensure(ctx context.Context, key string, value any) (exists bool, err error)

	// Delete deletes cache for a specific key.
	Delete(ctx context.Context, key string) error
}

type TTLCache interface {
	Cache

	// SetWithTTL caches data for a specific key with custom TTL.
	SetWithTTL(ctx context.Context, key string, value any, ttl time.Duration) error

	// SetWithTTLNX caches data for a specific key with custom TTL.
	// NX -- Only set the key if it does not already exist.
	// returns true if set.
	SetWithTTLNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)

	// EnsureWithTTL gets cached data or sets if data does not exist.
	// value: is setting and getting value. Use pointers to get the value.
	// ttl: custom TTL to set the value.
	// exist: is true when cache key existed before.
	EnsureWithTTL(ctx context.Context, key string, value any, ttl time.Duration) (exists bool, err error)
}

type SetCache interface {
	SAdd(ctx context.Context, key string, members ...any) (addedAmount int64, err error)
	SRem(ctx context.Context, key string, members ...any) (removedAmount int64, err error)
}

// GetCached gets from cache or loads from external source.
func GetCached[T any](
	ctx context.Context,
	cache TTLCache,
	key string,
	ttl time.Duration,
	get func(ctx context.Context) (T, error),
) (T, error) {
	empty := *new(T)
	var val T
	err := cache.Get(ctx, key, &val)
	if err == nil {
		return val, nil
	}
	if err != nil && err != ErrNotFound {
		return empty, err
	}
	// if not found
	val, err = get(ctx)
	if err != nil {
		return empty, err
	}
	err = cache.SetWithTTL(ctx, key, &val, ttl)
	if err != nil {
		return empty, err
	}
	return val, nil
}

// GetAndCache gets from external source and stores in cache.
func GetAndCache[T any](
	ctx context.Context,
	cache TTLCache,
	key string,
	ttl time.Duration,
	get func(ctx context.Context) (T, error),
) (T, error) {
	empty := *new(T)
	val, err := get(ctx)
	if err != nil {
		return empty, err
	}
	err = cache.SetWithTTL(ctx, key, &val, ttl)
	if err != nil {
		return empty, err
	}
	return val, nil
}
