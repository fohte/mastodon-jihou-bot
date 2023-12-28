package main

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		cfg := &Config{
			AccessToken: "test"
		}
		got, err := Run(ctx, NewMastodonService(cfg))
		want := "post success: https://social.fohte.net/@jihou/105029911841732872"

		if err != nil {
			t.Fatalf("want nil, but got %v", err)
		}

		if *got != want {
			t.Errorf("want %s, but got %s", want, *got)
		}

	})
}
