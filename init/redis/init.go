package redis

import (
	"github.com/redis/go-redis/v9"
)

// NewPool 初始化redis连接池
func NewPool() *redis.ClusterClient {
	opts := redis.ClusterOptions{
		MaxIdleConns: 10000,
		PoolSize:     12000, // max number of connections
		NewClient: func(opt *redis.Options) *redis.Client {
			opt.Addr = ":6379"
			return redis.NewClient(opt)
		},
	}
	return redis.NewClusterClient(&opts)
}
