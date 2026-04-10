package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001009_complementOfBase10Integer/go"
)

func TestBitwiseComplement(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 1},
		{"one", 1, 0},
		{"two", 2, 1},
		{"five", 5, 2},
		{"seven", 7, 0},
		{"ten", 10, 5},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.BitwiseComplement(tc.n)
			if got != tc.want {
				t.Fatalf("BitwiseComplement(%v) = %v expected: %v", tc.n, got, tc.want)
			}
		})
	}
}
