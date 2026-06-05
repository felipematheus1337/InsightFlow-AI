package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/felipematheus1337/InsightFlow-AI/go-consumer/interfaces"
)

type DocumentService struct {
	pdf     *PDFService
	storage interfaces.Storage
}

func NewDocumentService(pdf *PDFService, storage interfaces.Storage) *DocumentService {
	return &DocumentService{pdf: pdf, storage: storage}
}

func (s *DocumentService) GenerateAndStore(ctx context.Context, texto string) (string, error) {
	fileName, content, err := s.pdf.CreatePdf(texto)
	if err != nil {
		return "", fmt.Errorf("generating pdf: %w", err)
	}

	err = s.storage.Upload(ctx, fileName, bytes.NewReader(content), int64(len(content)))
	if err != nil {
		return "", fmt.Errorf("storing pdf: %w", err)
	}

	return fileName, nil
}
