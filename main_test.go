package main_test

import (
	"context"
	"testing"
	"time"

	"github.com/fohte/mastodon-jihou-bot"
	"github.com/fohte/mastodon-jihou-bot/mock_main"
	"github.com/google/go-cmp/cmp"
	"github.com/mattn/go-mastodon"
	"go.uber.org/mock/gomock"
)

func TestRun(t *testing.T) {
	cases := []struct {
		description string
		url         string
		now         time.Time
		content     string
	}{
		{
			description: "success",
			url:         "https://social.fohte.net/@jihou/105029911841732872",
			now:         time.Date(2023, 12, 28, 0, 0, 0, 0, time.UTC),
			content:     "📅 2023-12-28 (Thu)\n🕛 00:00\n\n📌 年末年始休暇: 残り 7/7 日 (100 %) [▓▓▓▓▓▓▓▓▓▓]",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			serviceMock := mock_main.NewMockMastodonService(ctrl)
			serviceMock.EXPECT().PostStatus(
				ctx,
				DiffEq(
					&mastodon.Toot{
						Status:     c.content,
						Visibility: "unlisted",
					}),
			).Return(&mastodon.Status{
				URL: c.url,
			}, nil)

			got, err := main.Run(
				ctx,
				serviceMock,
				NewTestTimeProvider(c.now),
			)

			if err != nil {
				t.Fatalf("want nil, but got %v", err)
			}

			want_message := "post success: " + c.url

			if diff := cmp.Diff(want_message, *got); diff != "" {
				t.Errorf("missmatch (-want +got):\n%s", diff)
			}
		})
	}
}
