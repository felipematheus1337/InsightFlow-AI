package messaging

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type MessageHandler func(ctx context.Context, payload []byte) error

type SQSService struct {
	subscriber *sqs.Subscriber
}

func NewSQSService(awsCfg aws.Config) (*SQSService, error) {

	logger := watermill.NewStdLogger(false, false)

	subscribe, err := sqs.NewSubscriber(sqs.SubscriberConfig{
		AWSConfig: awsCfg,
	}, logger)

	if err != nil {
		return nil, fmt.Errorf("error creating SQS Subscriber: %w", err)
	}

	return &SQSService{subscriber: subscribe}, nil
}

func (s *SQSService) Subscribe(ctx context.Context, queuename string, handler MessageHandler) error {
	messages, err := s.subscriber.Subscribe(ctx, queuename)

	if err != nil {
		return fmt.Errorf("error subscribing messages: %w", err)
	}
	for msg := range messages {
		if err := handler(ctx, msg.Payload); err != nil {
			msg.Nack()
			continue
		}

		msg.Ack()
	}

	return nil
}

func (s *SQSService) Close() error {
	return s.subscriber.Close()
}
