package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type MittenEvent struct {
	Word string
}

func HandleRequest(ctx context.Context, event MittenEvent) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := sqs.NewFromConfig(cfg)
	queueUrl := os.Getenv("QUEUE_URL")

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(event.Word),
		QueueUrl:    aws.String(queueUrl),
	}

	_, err = client.SendMessage(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to enqueue message: %v", err)
	}
	return fmt.Sprintf("sucessfully enqueued: %v", event.Word), nil
}

func main() {
	lambda.Start(HandleRequest)
}
