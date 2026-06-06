package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/config"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/handler"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/messaging"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	creds, err := cfg.AWS.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatalf("Error retrieving AWS credentials: %v", err)
	}

	storage, err := service.NewMinioStorage(
		cfg.S3Endpoint,
		creds.AccessKeyID,
		creds.SecretAccessKey,
		cfg.S3Bucket,
		cfg.S3UseSSL,
	)
	if err != nil {
		log.Fatalf("Error creating storage: %v", err)
	}

	documentService := service.NewDocumentService(&service.PDFService{}, storage)
	docHandler := handler.New(documentService)

	sqsService, err := messaging.NewSQSService(cfg.AWS)
	if err != nil {
		log.Fatalf("Error creating SQS service: %v", err)
	}
	defer sqsService.Close()
	log.Printf("consumindo a fila %q...", cfg.QueueName)
	if err := sqsService.Subscribe(ctx, cfg.QueueName, docHandler.Handle); err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	log.Println("consumer encerrado.")
}
