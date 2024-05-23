package repository

import (
	"bytes"
	"context"
	"fmt"
	"gin/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/skip2/go-qrcode"
)

type AWSRepository interface {
	CreateAndSaveQrCodeInS3(ctx context.Context, cnh *string) (string, error)
}

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

func (a *AWS) CreateAndSaveQrCodeInS3(ctx context.Context, cnh *string) (string, error) {

	qrCodeData := fmt.Sprintf("http://localhost:8080/api/v1/drivers/%s", *cnh)
	qrCode, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)

	if err != nil {
		return "Error creating qrcode", err
	}

	svc := s3.New(a.conn)

	fileName := fmt.Sprintf("qrcodes/%s.png", *cnh)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("venture-s3-bucket"),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(qrCode),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "Error while saving qrcode in aws", err
	}

	qrCodeURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "venture-s3-bucket", fileName)

	return qrCodeURL, nil

}
