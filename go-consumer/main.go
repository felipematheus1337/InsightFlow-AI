package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/config"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/handler"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/messaging"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load(ctx)

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	sqsService, err := messaging.NewSQSService(cfg.AWS)

	if err != nil {
		log.Fatalf("Error creating SQS service: %v", err)
	}

	defer sqsService.Close()

	if err := sqsService.Subscribe(ctx, cfg.QueueName, handler.Handle); err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

}
