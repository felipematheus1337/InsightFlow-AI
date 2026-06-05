package service

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client *minio.Client
	bucket string
}

func NewMinioStorage(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinioStorage, error) {

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("creating minio client: %w", err)
	}
	return &MinioStorage{client: client, bucket: bucket}, nil
}

func (m *MinioStorage) Upload(ctx context.Context,
	name string, r io.Reader, size int64) error {

	_, err := m.client.PutObject(ctx, m.bucket, name, r, size,
		minio.PutObjectOptions{ContentType: "application/pdf"})

	if err != nil {
		return fmt.Errorf("uploading %s: %w", name, err)
	}

	return nil
}
