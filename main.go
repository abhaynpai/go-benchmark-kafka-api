package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Payload is the structure used to accept Kafka Producer payload
type Payload struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Its working!")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", sayHello)

	e.Logger.Fatal(e.Start(":1234"))
}

func sayHello(c echo.Context) error {
	p := new(Payload)

	if err := c.Bind(p); err != nil {
		log.Println(err)
		return err
	}

	k, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	if err != nil {
		panic(err)
	}

	k.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(p.Message),
	}, nil)

	return c.JSON(200, p)
}
