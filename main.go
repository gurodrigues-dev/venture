package main

import (
	"gin/config"
	_ "gin/docs"
	"gin/internal/controllers"
	"gin/internal/repository"
	"gin/internal/service"
	"log"
)

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	repo, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("error creating repository: %s", err.Error())
	}

	aws, err := repository.NewAwsConnection()
	if err != nil {
		log.Fatalf("error creating aws connection: %s", err.Error())
	}

	redis, err := repository.NewRedisClient()
	if err != nil {
		log.Fatalf("error creating redis connection: %s", err.Error())
	}

	kafka, err := repository.NewKafkaClient()
	if err != nil {
		log.Fatalf("error creating kafka connection: %s", err.Error())
	}

	controller := controllers.New(
		service.NewService(repo, aws, redis, kafka),
		service.NewResponsibleService(repo),
		service.NewDriverService(repo),
		service.NewSchoolService(repo),
		service.NewChildService(repo),
	)

	log.Printf("initing service: %s", config.Name)
	controller.Start()

}
