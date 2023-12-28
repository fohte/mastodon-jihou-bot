package main_test

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

type cmpMatcher struct {
	expected interface{}
	diff     string
}

func (m *cmpMatcher) Matches(x interface{}) bool {
	m.diff = cmp.Diff(x, m.expected)
	return m.diff == ""
}

func (m *cmpMatcher) String() string {
	if m.diff == "" {
		return ""
	}
	return fmt.Sprintf("diff(-got +want): %s", m.diff)
}

func DiffEq(expected interface{}) gomock.Matcher {
	return &cmpMatcher{expected: expected}
}
