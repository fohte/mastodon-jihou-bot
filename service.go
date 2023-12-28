package main

import (
	"context"

	"github.com/mattn/go-mastodon"
)

type mastodonService struct {
	client *mastodon.Client
}

func NewMastodonService(cfg *Config) *mastodonService {
	client := mastodon.NewClient(&mastodon.Config{
		Server:      SERVER_URL,
		AccessToken: cfg.AccessToken,
	})

	return &mastodonService{
		client: client,
	}
}

func (s *mastodonService) PostStatus(ctx context.Context, toot *mastodon.Toot) (*mastodon.Status, error) {
	status, err := s.client.PostStatus(ctx, toot)
	if err != nil {
		return nil, err
	}

	return status, nil
}
