package configuration

import (
	"online-store-golang/errs"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedis(config Config) *redis.Client {
	host := config.Get("REDIS_HOST")
	port := config.Get("REDIS_PORT")
	maxPoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MAX_SIZE"))
	errs.PanicIfError(err)
	minIdlePoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MIN_IDLE_SIZE"))
	errs.PanicIfError(err)

	redisStore := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
	return redisStore
}
