package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/xxxxxxx_title/go"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		param    int
		expected int
	}{
		{"first case", 1, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.Solution(tc.param)
			if got != tc.expected {
				t.Fatalf("Solution(%d) = %d, expected %d", tc.param, got, tc.expected)
			}
		})
	}
}
