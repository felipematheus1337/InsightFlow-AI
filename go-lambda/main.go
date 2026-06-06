package go_lambda

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var (
	s3Client  *s3.Client
	sesClient *sesv2.Client

	sender    = os.Getenv("SENDER_EMAIL")
	recipient = os.Getenv("RECIPIENT_EMAIL")
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	s3Client = s3.NewFromConfig(cfg)
	sesClient = sesv2.NewFromConfig(cfg)
}

func handler(ctx context.Context, event events.S3Event) error {
	for _, record := range event.Records {
		bucket := record.S3.Bucket.Name

		key, err := url.QueryUnescape(record.S3.Object.Key)

		if err != nil {
			return fmt.Errorf("url unescape error: %w", err)
		}

		log.Printf("S3 Object Key: %s", key)

		pdfBytes, err := baixarObjeto(ctx, bucket, key)

		if err != nil {
			return fmt.Errorf("baixarObjeto error: %w", err)
		}

		if err := enviarEmailComAnexo(ctx, key, pdfBytes); err != nil {
			return fmt.Errorf("enviarEmailComAnexo error: %w", err)
		}

		log.Println("Email send sucessfully")
	}

	return nil
}

func baixarObjeto(ctx context.Context, bucket, key string) ([]byte, error) {
	out, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()

	return io.ReadAll(out.Body)
}

func enviarEmailComAnexo(ctx context.Context, key string, conteudo []byte) error {
	nomeArquivo := path.Base(key)

	_, err := sesClient.SendEmail(ctx, &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(sender),
		Destination: &types.Destination{
			ToAddresses: []string{recipient},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String("Novo PDF recebido: " + nomeArquivo),
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: aws.String("Segue em anexo o arquivo que chegou no bucket S3."),
					},
				},
				Attachments: []types.Attachment{
					{
						FileName:   aws.String(nomeArquivo),
						RawContent: conteudo,
					},
				},
			},
		},
	})
	return err
}

func main() {
	lambda.Start(handler)
}
