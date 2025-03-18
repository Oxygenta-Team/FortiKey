package redis

import "github.com/redis/go-redis/v9"

type Cacher interface {
	Pipeline() *Pipeliner
	Do(call func(client *redis.Client) *redis.StatusCmd) error
	Get(call func(client *redis.Client) (any, error)) (any, error)
}
