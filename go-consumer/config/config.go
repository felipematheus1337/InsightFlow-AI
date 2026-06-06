package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/joho/godotenv"
)

type Config struct {
	QueueName  string
	AWS        aws.Config
	S3Endpoint string
	S3Bucket   string
	S3UseSSL   bool
}

func Load(ctx context.Context) (*Config, error) {

	_ = godotenv.Load()
	var missing []string
	region := get("AWS_REGION", &missing)
	accessKey := get("AWS_ACCESS_KEY_ID", &missing)
	secretKey := get("AWS_SECRET_ACCESS_KEY", &missing)
	queueName := get("SQS_QUEUE_NAME", &missing)
	s3endpoint := get("S3_ENDPOINT", &missing)
	s3Bucket := get("S3_BUCKET", &missing)
	s3useSSLStr := get("S3_USE_SSL", &missing)

	useSSL, err := strconv.ParseBool(s3useSSLStr)

	if err != nil {
		return nil, err
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("loading AWS config: %w", err)
	}

	return &Config{QueueName: queueName, AWS: awsCfg, S3Endpoint: s3endpoint, S3Bucket: s3Bucket, S3UseSSL: useSSL}, nil
}

func get(key string, missing *[]string) string {
	v := os.Getenv(key)

	if v == "" {
		*missing = append(*missing, key)
	}

	return v
}
