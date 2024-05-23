package service

import (
	"context"
	"gin/internal/repository"
)

type RedisService struct {
	redisrepository repository.RedisRepository
}

func NewRedisService(redis repository.RedisRepository) *RedisService {
	return &RedisService{
		redisrepository: redis,
	}
}

func (rds *RedisService) SaveKeyAndValue(ctx context.Context) {

}

func (rds *RedisService) FindKeyRedis(ctx context.Context) {

}

func (rds *RedisService) IsSismember(ctx context.Context) {

}

func (rds *RedisService) VerifyToken(ctx context.Context) {

}
