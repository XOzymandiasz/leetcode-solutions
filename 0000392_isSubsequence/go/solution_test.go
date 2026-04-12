package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000392_isSubsequence/go"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "basic true", s: "abc", t: "ahbgdc", want: true},
		{name: "basic false", s: "axc", t: "ahbgdc", want: false},
		{name: "empty s", s: "", t: "ahbgdc", want: true},
		{name: "empty t", s: "abc", t: "", want: false},
		{name: "both empty", s: "", t: "", want: true},
		{name: "same strings", s: "abc", t: "abc", want: true},
		{name: "longer s than t", s: "abcd", t: "abc", want: false},
		{name: "single char true", s: "a", t: "a", want: true},
		{name: "single char false", s: "b", t: "a", want: false},
		{name: "repeated chars", s: "aaa", t: "aaabbb", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution.IsSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
