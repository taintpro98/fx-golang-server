package queue

import (
	"fmt"
	"fx-golang-server/config"
	"strconv"

	"github.com/hibiken/asynq"
)

func getConfig(cfgQueue config.RedisQueueConfig) *asynq.RedisClientOpt {
	redisAddr := fmt.Sprintf("%s:%s", cfgQueue.Host, cfgQueue.Port)

	dbQueue, err := strconv.Atoi(cfgQueue.DB)
	if err != nil {
		dbQueue = 0
	}

	clientOptions := asynq.RedisClientOpt{
		Addr:     redisAddr,
		DB:       dbQueue,
		Password: cfgQueue.Password,
		Username: cfgQueue.Username,
		// PoolSize:     cfgQueue.PoolSize,
		// WriteTimeout: cfgQueue.WriteTimeOut * time.Second,
		// ReadTimeout:  cfgQueue.ReadTimeOut * time.Second,
		// DialTimeout:  cfgQueue.DialTimeOut * time.Second,
	}
	return &clientOptions
}

func NewServer(cfg config.RedisQueueConfig) *asynq.Server {
	clientOptions := getConfig(cfg)

	return asynq.NewServer(
		clientOptions,
		asynq.Config{
			Concurrency: 20,
			Queues: map[string]int{
				"critical": 6,
				"default":  4,
			},
		},
	)
}

func NewClient(cfg config.RedisQueueConfig) *asynq.Client {
	clientOptions := getConfig(cfg)

	return asynq.NewClient(clientOptions)
}
