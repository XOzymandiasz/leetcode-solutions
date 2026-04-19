package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001207_uniqueNumberOfOccurences/go"
)

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want bool
	}{
		{"single number", []int{1, 1, 1, 1}, true},
		{"negative case", []int{1, 2, 2, 1}, false},
		{"positive case", []int{1, 1, 1, 2, 2, 3}, true},
		{"negative values", []int{-1, -1, -1, 2, 2, 3}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.UniqueOccurrences(tc.arr)
			if tc.want != got {
				t.Errorf("UniqueOccurrences(%v) = %v expected: %v", tc.arr, got, tc.want)
			}
		})
	}
}
