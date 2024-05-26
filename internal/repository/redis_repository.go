package repository

import "github.com/go-redis/redis"

type RedisRepository struct {
	conn *redis.Client
}

func NewRedisRepository(conn *redis.Client) *RedisRepository {
	return &RedisRepository{
		conn: conn,
	}
}
