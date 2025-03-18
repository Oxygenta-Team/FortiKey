package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Oxygenta-Team/FortiKey/pkg/db"
)

type Client struct {
	client *redis.Client
}

type Pipeliner struct {
	pipeliner redis.Pipeliner
}

func NewPipeliner(pipe redis.Pipeliner) *Pipeliner {
	return &Pipeliner{pipeliner: pipe}
}

func NewClient(config *db.Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Username: config.User,
		Password: config.Password,
		DB:       0,
	})

	if pong, err := rdb.Ping(context.Background()).Result(); err != nil && pong == "PONG" {
		return nil, err
	}

	return &Client{client: rdb}, nil
}

func (cl *Client) Pipeline() *Pipeliner {
	pipe := cl.client.Pipeline()
	return &Pipeliner{pipeliner: pipe}
}

func (cl *Client) Do(call func(client *redis.Client) *redis.StatusCmd) error {
	cmd := call(cl.client)
	return cmd.Err()
}

// TODO refactor
func (cl *Client) Get(call func(client *redis.Client) (any, error)) (any, error) {
	return call(cl.client)
}

func (pipe *Pipeliner) Add(call func(pipe redis.Pipeliner) *redis.StatusCmd) error {
	cmd := call(pipe.pipeliner)
	return cmd.Err()
}

func (pipe *Pipeliner) Do(ctx context.Context) ([]redis.Cmder, error) {
	return pipe.pipeliner.Exec(ctx)
}
