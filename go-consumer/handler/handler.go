package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/response"
	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/service"
)

type Handler struct {
	docService *service.DocumentService
}

func New(docService *service.DocumentService) *Handler {
	return &Handler{docService: docService}
}

func (h *Handler) Handle(ctx context.Context, payload []byte) error {
	var resp response.Response
	if err := json.Unmarshal(payload, &resp); err != nil {
		return fmt.Errorf("unmarshaling payload: %w", err)
	}

	log.Printf("processando relatório (%d chars)", len(resp.Relatorio))

	fileName, err := h.docService.GenerateAndStore(ctx, resp.Relatorio)
	if err != nil {
		return fmt.Errorf("generating/storing document: %w", err)
	}

	log.Printf("relatório armazenado: %s", fileName)
	return nil
}
