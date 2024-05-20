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

	responsible := service.NewResponsibleService(repo)
	driver := service.NewDriverService(repo)
	service := service.NewService(repo, aws, redis, kafka)
	controller := controllers.New(service, responsible, driver)

	log.Printf("initing service: %s", config.Name)
	controller.Start()

}
