package rediscache

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/gob"
	"errors"
	"github.com/redis/go-redis/v9"
	cache "komek/pkg/redis_cache"
	"time"
)

const (
	pingTimeout  = 15 * time.Second
	txMaxRetries = 1000 // transaction max amount of retries
)

var (
	ErrTxMaxRetries = errors.New("increment reached maximum number of retries")
)

// Config configures redis cache
type Config struct {
	Addr       string
	User       string
	Pass       string
	DB         int
	TLSEnabled bool
}

// Client is a wrapper on client interfaces.
// Used for better testing.
type Client interface {
	redis.Cmdable
	redis.UniversalClient
}

// redisPipeliner is wrapper on redis pipeliner.
// Pipeliner is used to create redis pipes (https://redis.io/docs/manual/pipelining/).
// In this project it is used for better testing.
type redisPipeliner interface {
	redis.Pipeliner
}

// RedisCache is Redis Client that stores data in JSON.
type RedisCache struct {
	rdb Client
}

// NewWithConfig creates new instance of redis cache
func NewWithConfig(cfg Config) (*RedisCache, error) {
	opts := &redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.User,
		Password: cfg.Pass,
		DB:       cfg.DB,
	}
	if cfg.TLSEnabled {
		opts.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	c := &RedisCache{
		rdb: redis.NewClient(opts),
	}
	// ping to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()
	cmd := c.rdb.Ping(ctx)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return c, nil
}

func NewWithClient(client Client) *RedisCache {
	return &RedisCache{
		rdb: client,
	}
}

// Exists checks if cache key exists.
func (c RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	cmd := c.rdb.Exists(ctx, key)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	if cmd.Val() == 1 {
		return true, nil
	}
	return false, nil
}

// get gets cached data for a specific key for any redis command type.
func get(ctx context.Context, cmdable redis.Cmdable, key string, value any) error {
	cmd := cmdable.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return cache.ErrNotFound
	} else if cmd.Err() != nil {
		return cmd.Err()
	}
	return decode([]byte(cmd.Val()), value)
}

// Get gets cached data for a specific key.
func (c RedisCache) Get(ctx context.Context, key string, value any) error {
	return get(ctx, c.rdb, key, value)
}

// setWithTTL caches data for a specific key with custom TTL for any redis command type.
func setWithTTL(ctx context.Context, cmdable redis.Cmdable, key string, value any, ttl time.Duration) error {
	val, err := encode(value)
	if err != nil {
		return err
	}
	return cmdable.Set(ctx, key, string(val), ttl).Err()
}

// setWithTTL caches data for a specific key with custom TTL for any redis command type.
// NX -- Only set the key if it does not already exist.
// returns true if set.
func setWithTTLNX(ctx context.Context, cmdable redis.Cmdable, key string, value any, ttl time.Duration) (bool, error) {
	val, err := encode(value)
	if err != nil {
		return false, err
	}
	return cmdable.SetNX(ctx, key, string(val), ttl).Result()
}

// Set caches data for a specific key with timeout.
func (c RedisCache) Set(ctx context.Context, key string, value any) error {
	return setWithTTL(ctx, c.rdb, key, value, 0)
}

// SetNX caches data for a specific key with a timeout.
// NX -- Only set the key if it does not already exist.
// returns true if set.
func (c RedisCache) SetNX(ctx context.Context, key string, value any) (bool, error) {
	return setWithTTLNX(ctx, c.rdb, key, value, 0)
}

// SetWithTTL caches data for a specific key with custom TTL.
func (c RedisCache) SetWithTTL(ctx context.Context, key string, value any, ttl time.Duration) error {
	return setWithTTL(ctx, c.rdb, key, value, ttl)
}

// SetWithTTLNX caches data for a specific key with custom TTL.
// NX -- Only set the key if it does not already exist.
// returns true if set.
func (c RedisCache) SetWithTTLNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	return setWithTTLNX(ctx, c.rdb, key, value, ttl)
}

// ensureTransaction creates transaction function for EnsureWithTTL method.
// returns `func(cmdable redis.Cmdable) error`, that is testable by mock tools.
func ensureTransaction(ctx context.Context, key string, value any, ttl time.Duration, exists *bool) func(cmdable redis.Cmdable) error {
	return func(cmdable redis.Cmdable) error {
		// init exists pointer to avoid panics
		if exists == nil {
			exists = &[]bool{true}[0]
		}
		*exists = false
		// get current key value and assign it
		err := get(ctx, cmdable, key, value)
		if err != nil && err != cache.ErrNotFound {
			return err
		} else if err == nil {
			*exists = true
			return nil
		}

		// if key is empty, set new value
		_, err = cmdable.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			return setWithTTL(ctx, pipe, key, value, ttl)
		})
		return err
	}
}

// execTx watches key and executes transaction
func (c RedisCache) execTx(
	ctx context.Context,
	key string,
	txf func(tx *redis.Tx) error,
) error {
	// Retry if the key has been changed.
	for i := 0; i < txMaxRetries; i++ {
		// wrap txf to apply Watch()
		err := c.rdb.Watch(ctx, txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return err
	}

	return ErrTxMaxRetries
}

// EnsureWithTTL gets cached data or sets if data does not exist.
// value: is setting and getting value. Use pointers to get the value.
// ttl: custom TTL to set the value.
// exist: is true when cache key existed before.
func (c RedisCache) EnsureWithTTL(ctx context.Context, key string, value any, ttl time.Duration) (exists bool, err error) {
	// transactional function.
	txf := ensureTransaction(ctx, key, value, ttl, &exists)

	// execute transaction
	err = c.execTx(ctx, key, func(tx *redis.Tx) error { return txf(tx) })
	if err != nil {
		return false, nil
	}
	return exists, err
}

// Ensure gets cached data or sets if data does not exist.
// value: is setting and getting value. Use pointers to get the value.
// exist: is true when cache key existed before.
func (c RedisCache) Ensure(ctx context.Context, key string, value any) (exists bool, err error) {
	return c.EnsureWithTTL(ctx, key, value, 0)
}

// Delete deletes cache for a specific key.
func (c RedisCache) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c RedisCache) SAdd(ctx context.Context, key string, members ...any) (addedAmount int64, err error) {
	return c.rdb.SAdd(ctx, key, members...).Result()
}

func (c RedisCache) SRem(ctx context.Context, key string, members ...any) (removedAmount int64, err error) {
	return c.rdb.SRem(ctx, key, members...).Result()

}

func encode(val any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(val)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, val any) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(val)
}
