package redis_mock

import (
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"

	r "github.com/Oxygenta-Team/FortiKey/pkg/db/redis"
)

type MockRedisClient struct {
	client *redis.Client
}

func NewMockRedisClient() (*MockRedisClient, redismock.ClientMock) {
	db, mock := redismock.NewClientMock()
	return &MockRedisClient{
		client: db,
	}, mock
}

func (cl *MockRedisClient) Pipeline() *r.Pipeliner {
	pipe := cl.client.Pipeline()
	return r.NewPipeliner(pipe)
}

func (cl *MockRedisClient) Do(call func(client *redis.Client) *redis.StatusCmd) error {
	cmd := call(cl.client)
	return cmd.Err()
}

// TODO refactor
func (cl *MockRedisClient) Get(call func(client *redis.Client) (any, error)) (any, error) {
	return call(cl.client)
}
