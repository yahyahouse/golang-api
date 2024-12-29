package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"exam/config"
	redisClient "exam/internal/delivery/redis"
)

type RedisRepository interface {
	GetAllKey(ctx context.Context) (value []string, err error)
	GetKey(ctx context.Context, key string) (value string, err error)
	SetKey(ctx context.Context, key string, value interface{}) (err error)
	DeleteKey(ctx context.Context, key string) (err error)
}

type redisRepository struct {
	redis *redisClient.Redis
	conf  *config.RedisConfig
}

const serviceName = "templateservice-bs"

func NewRedisRepository(redis *redisClient.Redis, cfg *config.RedisConfig) RedisRepository {
	return &redisRepository{redis, cfg}
}

func (repo redisRepository) GetAllKey(ctx context.Context) (value []string, err error) {
	start := time.Now()
	defer func() {
		log.Println("GetAllKey_REDIS_BACKEND", serviceName, err == nil, time.Duration(time.Since(start)))
	}()

	if !config.GetConfig().IsByPassDBRedis {
		var cursor uint64
		var keys []string
		for {
			keys, cursor, err = repo.redis.RedisClient.Scan(ctx, cursor, "*template|*", 0).Result()
			if err != nil {
				return value, err
			}
			value = append(value, keys...)
			if cursor == 0 {
				break
			}
		}
	}

	return value, nil
}

func (repo redisRepository) GetKey(ctx context.Context, key string) (value string, err error) {
	start := time.Now()
	defer func() {
		log.Println("GetKey_REDIS_BACKEND", serviceName, err == nil, time.Duration(time.Since(start)))
	}()

	if !config.GetConfig().IsByPassDBRedis {
		value, err = repo.redis.RedisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			err = fmt.Errorf("key doesn't exist : %s, data type : %T", redis.Nil, redis.Nil)
			return value, err
		} else if err != nil {
			return value, err
		}
	} else {
		return "Terima kasih atas pembayaran tagihan Halo senilai [price] melalui Tcops partial payment. Silakan restart ponsel Anda jika layanan belum aktif", nil
	}

	return value, nil
}

func (repo redisRepository) SetKey(ctx context.Context, key string, value interface{}) (err error) {
	start := time.Now()
	defer func() {
		log.Println("SetKey_REDIS_BACKEND", serviceName, err == nil, time.Duration(time.Since(start)))
	}()

	if !config.GetConfig().IsByPassDBRedis {
		err = repo.redis.RedisClient.Set(ctx, key, value, repo.conf.TimeToLive).Err()
	}

	return err
}

func (repo redisRepository) DeleteKey(ctx context.Context, key string) (err error) {
	start := time.Now()
	defer func() {
		log.Println("DeleteKey_REDIS_BACKEND", serviceName, err == nil, time.Duration(time.Since(start)))
	}()

	if !config.GetConfig().IsByPassDBRedis {
		if err := repo.redis.RedisClient.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	return nil
}
