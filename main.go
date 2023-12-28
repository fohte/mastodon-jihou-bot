package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v10"
	"github.com/mattn/go-mastodon"
)

type Config struct {
	AccessToken string `env:"MASTODON_ACCESS_TOKEN"`
}

const (
	SERVER_URL = "https://social.fohte.net"
)

type Event struct{}

func Run(ctx context.Context, service *mastodonService) (*string, error) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	timeReport := NewTimeReport(location)

	content := timeReport.CreateTimeReport()

	status, err := service.PostStatus(ctx, &mastodon.Toot{
		Status:     content,
		Visibility: "unlisted",
	})

	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("post success: %s", status.URL)
	return &message, nil
}

func handleRequest(ctx context.Context, event *Event) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	service := NewMastodonService(cfg)

	return Run(ctx, service)
}

func main() {
	lambda.Start(handleRequest)
}
