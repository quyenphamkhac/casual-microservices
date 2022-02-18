package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/transport"
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

	httpServer := transport.NewHttpServer(cfg)
	httpServer.Run(":3000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("signal notify: %v", sig)
}
