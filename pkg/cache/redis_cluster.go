package cache

import (
	"context"
	"fmt"
	"fx-golang-server/config"

	"github.com/redis/go-redis/v9"
)

type redisCluster struct {
	client *redis.ClusterClient
	cfg    config.RedisConfig
}

// CloseConnection implements IRedisClient.
func (r *redisCluster) CloseConnection() error {
	panic("unimplemented")
}

// Get implements IRedisClient.
func (r *redisCluster) Get(ctx context.Context, key string, outputType interface{}) error {
	panic("unimplemented")
}

// LPop implements IRedisClient.
func (r *redisCluster) LPop(ctx context.Context, queueName string, outputType interface{}) error {
	panic("unimplemented")
}

// Publish implements IRedisClient.
func (r *redisCluster) Publish(ctx context.Context, channel string, value interface{}) error {
	panic("unimplemented")
}

// RPush implements IRedisClient.
func (r *redisCluster) RPush(ctx context.Context, queueName string, value interface{}) error {
	panic("unimplemented")
}

// Set implements IRedisClient.
func (r *redisCluster) Set(ctx context.Context, key string, value interface{}, ttl int64) error {
	panic("unimplemented")
}

// Subscribe implements IRedisClient.
func (r *redisCluster) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	panic("unimplemented")
}

func NewRedisCluster(cfg config.RedisConfig) (IRedisClient, error) {
	redisOpt := &redis.ClusterOptions{
		Addrs:    []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		Password: cfg.Password,
		Username: cfg.Username,
	}

	client := redis.NewClusterClient(redisOpt)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &redisCluster{
		client: client,
		cfg:    cfg,
	}, err
}
