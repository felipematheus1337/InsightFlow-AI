package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/config"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/response"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/service"
)

func Handle(ctx context.Context, payload []byte) error {

	var response *response.Response
	if err := json.Unmarshal(payload, &response); err != nil {
		return err
	}

	log.Printf("Received a response: %s", response.Relatorio)

	cfg, err := config.Load(ctx)

	creds, err := cfg.AWS.Credentials.Retrieve(ctx)
	if err != nil {
		return err
	}

	var pdfService service.PDFService

	minioStorage, err := service.NewMinioStorage(
		"mockEndpoint",
		creds.AccessKeyID,
		creds.SecretAccessKey,
		"mockBucket",
		false,
	)

	if err != nil {
		return err
	}

	documentService := service.NewDocumentService(&pdfService, minioStorage)

	go documentService.GenerateAndStore(ctx, response.Relatorio)

	return nil
}
