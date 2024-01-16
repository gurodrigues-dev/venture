package utils

import (
	"gin/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func VeryifyEmailInAwsSes(emailAddress string) (bool, error) {

	config.LoadEnvironmentVariables()

	var (
		region       = config.GetRegionAws()
		awsAccessKey = config.GetAwsAccessKey()
		awsSecretKey = config.GetAwsSecretKey()
		awsTokenKey  = config.GetAwsTokenKey()
	)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			awsAccessKey,
			awsSecretKey,
			awsTokenKey),
	})

	if err != nil {
		return false, err
	}

	svc := ses.New(sess)

	verifyEmailInput := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String(emailAddress),
	}

	_, err = svc.VerifyEmailIdentity(verifyEmailInput)

	return true, nil
}

func DeleteEmailFromAwsSes(emailAddress string) (bool, error) {

	config.LoadEnvironmentVariables()

	var (
		region       = config.GetRegionAws()
		awsAccessKey = config.GetAwsAccessKey()
		awsSecretKey = config.GetAwsSecretKey()
		awsTokenKey  = config.GetAwsTokenKey()
	)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			awsAccessKey,
			awsSecretKey,
			awsTokenKey),
	})

	if err != nil {
		return false, err
	}

	svc := ses.New(sess)

	_, err = svc.DeleteVerifiedEmailAddress(&ses.DeleteVerifiedEmailAddressInput{
		EmailAddress: aws.String(emailAddress),
	})

	if err != nil {
		return false, err
	}

	return true, nil

}
