package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mattn/go-mastodon"
)

const (
	SERVER_URL = "https://social.fohte.net"
)

func init() {
	if os.Getenv("MASTODON_ACCESS_TOKEN") == "" {
		panic("MASTODON_ACCESS_TOKEN is not set")
	}
}

type Event struct{}

func handleRequest(ctx context.Context, event *Event) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	timeReport := NewTimeReport(location)

	content := timeReport.CreateTimeReport()

	client := mastodon.NewClient(&mastodon.Config{
		Server:      SERVER_URL,
		AccessToken: os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	status, err := client.PostStatus(ctx, &mastodon.Toot{
		Status:     content,
		Visibility: "unlisted",
	})

	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("post success: %s", status.URL)
	return &message, nil
}

func main() {
	lambda.Start(handleRequest)
}
