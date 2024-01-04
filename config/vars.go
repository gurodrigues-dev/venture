package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() (bool, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading file .env")

		return false, err
	}

	return true, nil
}

func GetHostDatabase() string {
	return os.Getenv("host")
}

func GetNameDatabase() string {
	return os.Getenv("database")
}

func GetPasswordDatabase() string {
	return os.Getenv("password")
}

func GetPortDatabase() string {
	return os.Getenv("port")
}

func GetUserDatabase() string {
	return os.Getenv("user")
}

func GetRegionAws() string {
	return os.Getenv("region")
}

func GetS3BucketAws() string {
	return os.Getenv("bucket")
}

func GetAwsAccessKey() string {
	return os.Getenv("accesskeyaws")
}

func GetAwsSecretKey() string {
	return os.Getenv("secretkeyaws")
}

func GetAwsTokenKey() string {
	return os.Getenv("tokenkeyaws")
}

func GetSecretKeyApi() string {
	return os.Getenv("secretkeyapi")
}

func GetRedisAddress() string {
	return os.Getenv("redisaddress")
}

func GetRedisPassword() string {
	return os.Getenv("redispassword")
}

func GetAwsEmailSource() string {
	return os.Getenv("emailawssource")
}
