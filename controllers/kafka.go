package controllers

import (
	"context"
	"encoding/json"
	"gin/models"
	"gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func ConsumerToCreateUsers(c *gin.Context) {

	kafkaConfig := models.KafkaConfig{
		BrokerAddress: "localhost:9092",
		Topic:         "kafka.create.users",
		GroupID:       "console-consumer-50786",
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaConfig.BrokerAddress},
		Topic:     kafkaConfig.Topic,
		GroupID:   kafkaConfig.GroupID,
		MaxBytes:  10e6, //10 MB
		Partition: 0,
	})

	defer r.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "reading messages w/ success",
	})

	for {

		m, err := r.ReadMessage(context.Background())

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error when reading messages in kafka",
				"error":   err.Error(),
			})

			return
		}

		var userData models.CreateUser

		err = json.Unmarshal(m.Value, &userData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error when unmarshal json",
				"error":   err.Error(),
			})
			return
		}

		_, err = repository.SaveUser(&userData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error when inserting into database",
				"error":   err.Error(),
			})
			return
		}

	}

}
