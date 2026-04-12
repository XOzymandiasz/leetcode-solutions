package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000011_containerWithMostWater/go"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name   string
		height []int
		expect int
	}{
		{"First Example", []int{2, 7, 4, 1, 6, 3, 9, 2, 5}, 35},
		{"Second Example", []int{1, 1}, 1},
		{"All same heights", []int{5, 5, 5, 5, 5}, 20},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MaxArea(tc.height)
			if got != tc.expect {
				t.Fatalf("MaxArea(%v) = %v expected: %v", tc.height, got, tc.expect)
			}
		})
	}
}
