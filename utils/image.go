package utils

import (
	"bytes"
	"fmt"
	"gin/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/skip2/go-qrcode"
)

func SaveQRCodeOfUser(cpf string) (string, error) {

	config.LoadEnvironmentVariables()

	var (
		region       = config.GetRegionAws()
		bucketName   = config.GetS3BucketAws()
		awsAccessKey = config.GetAwsAccessKey()
		awsSecretKey = config.GetAwsSecretKey()
		awsTokenKey  = config.GetAwsTokenKey()
	)

	qrCodeData := fmt.Sprintf("http://localhost:8080/api/v1/users/%s", cpf)
	qrCode, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)

	if err != nil {
		return "Error creating qrcode", err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, awsTokenKey),
	})

	if err != nil {
		return "Error while creating session aws", err
	}

	svc := s3.New(sess)

	fileName := fmt.Sprintf("qrcodes/%s.png", cpf)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(qrCode),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "Error while saving qrcode in aws", err
	}

	qrCodeURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName)

	return qrCodeURL, nil

}
