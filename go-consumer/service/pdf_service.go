package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"seehuhn.de/go/pdf"
)

type PDFService struct {
}

func (p *PDFService) CreatePdf(texto string) {
	uid, err := uuid.NewUUID()

	var fileName = fmt.Sprintf(uid.String() + time.Now().String() + ".pdf")

	if err != nil {
		panic(err)
	}

	w, err := pdf.Create(fileName, 9, nil)

	if err != nil {
		panic(err)
	}

	err = w.Close()

	if err != nil {
		panic(err)
	}
}
