package cache

import (
	"github.com/redis/go-redis/v9"
)

var ErrKeyNotExists = redis.Nil

type RedisCache struct {
	cmd redis.Cmdable
}

type Cache interface {
	//Set(ctx context.Context, question []domai) error
	//Get(ctx context.Context) ([]domain, error)
}

func NewCardRedisCache(cmd redis.Cmdable) Cache {
	return &RedisCache{cmd: cmd}
}
