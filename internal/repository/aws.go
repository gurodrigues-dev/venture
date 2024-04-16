package repository

import (
	"context"
	"gin/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWS struct {
	conn *session.Session
}

func NewAwsConnection() (*AWS, error) {

	conf := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(conf.Cloud.AccessKey, conf.Cloud.SecretKey, conf.Cloud.Token),
	})

	if err != nil {
		return nil, err
	}

	repo := &AWS{
		conn: sess,
	}

	return repo, nil

}

func (a *AWS) SaveImageBucket(ctx context.Context) {

}
