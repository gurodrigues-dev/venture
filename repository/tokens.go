package repository

import (
	"gin/config"
	"time"

	"github.com/go-redis/redis"
)

func SaveTokenToRedis(email, token string) (bool, error) {

	config.LoadEnvironmentVariables()

	var (
		redisAddress  = config.GetRedisAddress()
		redisPassword = config.GetRedisPassword()
	)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})

	err := client.Set(email, token, 10*time.Minute).Err()

	if err != nil {
		return false, err
	}

	return true, err

}
