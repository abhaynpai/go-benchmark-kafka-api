package main

import (
	"fmt"
	"log"

	"github.com/abhaynpai/go-test-kafka.git/producer"
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

	err := producer.InitKafka("localhost:9092")

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/produce", produce)

	e.Logger.Fatal(e.Start(":1234"))
}

func produce(c echo.Context) error {
	p := new(Payload)

	if err := c.Bind(p); err != nil {
		log.Println(err)
		c.Error(err)
		return err
	}

	err := producer.Produce(p.Topic, p.Message)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(200, p)
}
