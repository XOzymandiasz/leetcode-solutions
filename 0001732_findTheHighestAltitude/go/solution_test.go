package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001732_findTheHighestAltitude/go"
)

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		name string
		gain []int
		want int
	}{
		{name: "first example", gain: []int{-5, 1, 5, 0, -7}, want: 1},
		{name: "all positive", gain: []int{1, 2, 3}, want: 6},
		{name: "all negative", gain: []int{-1, -2, -3}, want: 0},
		{name: "mixed", gain: []int{3, -2, 5, -1}, want: 6},
		{name: "empty", gain: []int{}, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.LargestAltitude(tc.gain)
			if got != tc.want {
				t.Errorf("largestAltitude(%v) = %d; want %d", tc.gain, got, tc.want)
			}
		})
	}
}
