package service

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

type CmpMatcher struct {
	expected interface{}
	opts     cmp.Options
}

// Matches implementation for [gomock.Matcher]
func (cm *CmpMatcher) Matches(x interface{}) bool {
	if diff := cmp.Diff(cm.expected, x, cm.opts...); diff != "" {
		fmt.Printf(`Diff: (-want +got)
%s`, diff)
		return false
	}
	return true
}

// String implementation for [gomock.Matcher]
func (cm *CmpMatcher) String() string {
	return fmt.Sprintf("%+v (%T)", cm.expected, cm.expected)
}

// 等価演算に `github.com/google/go-cmp/cmp` を使用した [gomock.Matcher] implementation を返す
func NewCmpMatcher(expected interface{}, opts ...cmp.Option) gomock.Matcher {
	return &CmpMatcher{
		expected: expected,
		opts:     opts,
	}
}
