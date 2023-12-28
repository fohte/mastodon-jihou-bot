package main_test

import (
	"time"
)

type TestTimeProvider struct {
	currentTime time.Time
}

func NewTestTimeProvider(currentTime time.Time) *TestTimeProvider {
	return &TestTimeProvider{
		currentTime: currentTime,
	}
}

func (t *TestTimeProvider) Now() time.Time {
	return t.currentTime
}
