package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000032_longestValidParentheses/go"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"base", "(()", 2},
		{"base2", "()(()", 2},
		{"base3", "())()", 2},
		{"base4", "()(()()", 4},
		{"base5", "()(())", 6},
		{"base6", "()((()))", 8},
		{"base7", "()(((()))", 6},
		{"base8", ")()())", 4},
		{"base9", "(()())", 6},
		{"empty", "", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.LongestValidParentheses(tc.s)
			if got != tc.want {
				t.Fatalf("LongestValidParentheses(%v) = %v expected: %v", tc.s, tc.want, got)
			}
		})
	}
}
