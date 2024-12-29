package provider

import (
	configImport "exam/config"
	"exam/internal/delivery/redis"
	"exam/internal/repository"
)

var (
	redisConnection *redis.Redis
	repositoryRedis repository.RedisRepository
)

func RedisConnection(redis *redis.Redis) {
	if redis != nil {
		redisConnection = redis
	}
}

func getRedisConnection() *redis.Redis {
	if redisConnection != nil {
		return redisConnection
	}

	redisConfig := configImport.GetConfig().RedisConfig
	redisConnection := redis.InitRedis(redisConfig)
	return redisConnection
}

func getRepositoryRedis() repository.RedisRepository {
	if repositoryRedis == nil {
		cfg := configImport.GetConfig().RedisConfig
		repositoryRedis = repository.NewRedisRepository(getRedisConnection(), &cfg)
	}

	return repositoryRedis
}
