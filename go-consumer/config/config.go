package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/joho/godotenv"
)

type Config struct {
	QueueName string
	AWS       aws.Config
}

func Load(ctx context.Context) (*Config, error) {

	_ = godotenv.Load()
	var missing []string
	region := get("AWS_REGION", &missing)
	accessKey := get("AWS_ACCESS_KEY_ID", &missing)
	secretKey := get("AWS_SECRET_ACCESS_KEY", &missing)
	queueName := get("SQS_QUEUE_NAME", &missing)

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

	return &Config{QueueName: queueName, AWS: awsCfg}, nil
}

func get(key string, missing *[]string) string {
	v := os.Getenv(key)

	if v == "" {
		*missing = append(*missing, key)
	}

	return v
}
