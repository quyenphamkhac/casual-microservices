package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/transport"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load(".dev.env")
	if err != nil {
		log.Fatal("error loading env file")
	}

	cfg, err := config.NewServiceConfig()
	if err != nil {
		panic(err)
	}
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	httpServer := transport.NewHttpServer(cfg)
	httpServer.Run(":3000", ch)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("signal notify: %v", sig)
}
