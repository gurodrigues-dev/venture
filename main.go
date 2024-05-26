package main

import (
	"database/sql"
	"fmt"
	"gin/config"
	_ "gin/docs"
	"gin/internal/controllers"
	"gin/internal/repository"
	"gin/internal/service"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	gin.DisableConsoleColor()

	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	awsRepository := repository.NewAWSRepository(sess)

	driverRepository := repository.NewDriverRepository(db)
	driverService := service.NewDriverService(driverRepository, *awsRepository)
	driverController := controllers.NewDriverController(driverService)

	schoolRepository := repository.NewSchoolRepository(db)
	schoolService := service.NewSchoolService(schoolRepository)
	schoolController := controllers.NewSchoolController(schoolService, driverService)

	responsibleRepository := repository.NewResponsibleRepository(db)
	responsibleService := service.NewResponsibleService(responsibleRepository)
	responsibleController := controllers.NewResponsibleController(responsibleService)

	router := gin.Default()
	responsibleController.RegisterRoutes(router)
	driverController.RegisterRoutes(router)
	schoolController.RegisterRoutes(router)
	log.Printf("initing service: %s", config.Name)
	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
