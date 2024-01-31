package redis

type Option func(*Redis)

func Some(size int) Option {
	return func(r *Redis) {
	}
}
