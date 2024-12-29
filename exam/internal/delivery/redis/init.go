package redis

import (
	"exam/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisClient *redis.Client
}

func InitRedis(redisConfig config.RedisConfig) *Redis {
	redisClient := &Redis{
		RedisClient: redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:       redisConfig.MasterName,
			SentinelAddrs:    redisConfig.SentinelAddrs,
			SentinelUsername: redisConfig.User,
			SentinelPassword: redisConfig.Password,
			DB:               redisConfig.Database,
			PoolSize:         redisConfig.PoolSize,
			DialTimeout:      redisConfig.Timeout,
			MaxIdleConns:     redisConfig.MaxIdleConns,
		}),
	}

	return redisClient
}

func (s *Redis) GetDBCon() *redis.Client {
	return s.RedisClient
}
