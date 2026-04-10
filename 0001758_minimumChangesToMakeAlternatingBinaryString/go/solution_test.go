package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001758_minimumChangesToMakeAlternatingBinaryString/go"
)

func TestMinOperations(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{"zero", "0", 0},
		{"one", "1", 0},
		{"alternating", "1010", 0},
		{"first example", "0010", 1},
		{"second example", "0010100", 2},
		{"same characters", "0000", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MinOperations(tc.s)

			if got != tc.want {
				t.Fatalf("minOperations(%v) = %v expected: %v", tc.s, got, tc.want)
			}
		})
	}
}
