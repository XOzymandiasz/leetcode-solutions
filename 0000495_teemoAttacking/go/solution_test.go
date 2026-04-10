package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000495_teemoAttacking/go"
)

func TestFindPoisonedDuration(t *testing.T) {
	testCases := []struct {
		name     string
		time     []int
		duration int
		want     int
	}{
		{"first", []int{1, 4, 6}, 3, 8},
		{"second", []int{1, 6, 7}, 8, 14},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.FindPoisonedDuration(tc.time, tc.duration)
			if got != tc.want {
				t.Fatalf("given: time = %v duration = %v expected: %v, got: %v", tc.time, tc.duration, tc.want, got)
			}
		})
	}
}
