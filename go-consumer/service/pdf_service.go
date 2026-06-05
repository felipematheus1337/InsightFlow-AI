package service

import (
	"bytes"
	"fmt"
	"io"

	"github.com/google/uuid"
	"seehuhn.de/go/pdf"
	"seehuhn.de/go/pdf/document"
	"seehuhn.de/go/pdf/font/standard"
)

type PDFService struct {
}

func (p *PDFService) CreatePdf(texto string) (fileName string, content []byte, err error) {

	fileName = uuid.NewString() + ".pdf"
	var buf bytes.Buffer

	if err := writePDF(&buf, texto); err != nil {
		return "", nil, fmt.Errorf("failed to write pdf: %w", err)
	}
	return fileName, buf.Bytes(), nil
}

func writePDF(w io.Writer, texto string) error {
	doc, err := document.WriteSinglePage(w, document.A4, pdf.V2_0, nil)
	if err != nil {
		return fmt.Errorf("creating page: %w", err)
	}
	font, err := standard.Helvetica.New()
	if err != nil {
		return fmt.Errorf("loading font: %w", err)
	}
	const (
		fontSize = 12.0
		leading  = 16.0
		marginX  = 50.0
		startY   = 800.0
	)

	doc.TextSetFont(font, fontSize)
	doc.TextBegin()
	doc.TextFirstLine(marginX, startY)

	for i, linha := range bytes.Split([]byte(texto), []byte("\n")) {
		if i > 0 {
			doc.TextSecondLine(0, -leading)
		}
		doc.TextShow(string(linha))
	}

	doc.TextEnd()
	
	return doc.Close()
}
