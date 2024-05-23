package service

import (
	"context"
	"gin/internal/repository"
)

type AWSService struct {
	awsrepository repository.AWSRepository
}

func NewAWSService(aws repository.AWSRepository) *AWSService {
	return &AWSService{
		awsrepository: aws,
	}
}

func (as *AWSService) CreateAndSaveQrCodeInS3(ctx context.Context, cnh *string) (string, error) {
	return as.awsrepository.CreateAndSaveQrCodeInS3(ctx, cnh)
}
