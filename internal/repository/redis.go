package repository

import (
	"context"
	"gin/config"

	"github.com/go-redis/redis"
)

type RedisRepository interface {
	SaveKeyAndValue(ctx context.Context)
	FindKeyRedis(ctx context.Context)
	IsSismember(ctx context.Context)
	VerifyToken(ctx context.Context)
}

type Redis struct {
	conn *redis.Client
}

func NewRedisClient() (*Redis, error) {

	conf := config.Get()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Address,
		Password: conf.Cache.Password,
		DB:       0,
	})

	repo := &Redis{
		conn: client,
	}

	return repo, nil

}

func (r *Redis) SaveKeyAndValue(ctx context.Context) {

}

func (r *Redis) FindKeyRedis(ctx context.Context) {

}

func (r *Redis) IsSismember(ctx context.Context) {

}

func (r *Redis) VerifyToken(ctx context.Context) {

}
