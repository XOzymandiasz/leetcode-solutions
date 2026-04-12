package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001679_maxNumberOfK-SumPairs/go"
)

var testCases = []struct {
	name   string
	nums   []int
	k      int
	expect int
}{
	{"first case", []int{1, 2, 3, 4}, 5, 2},
	{"second case", []int{3, 1, 3, 4, 3}, 6, 1},
	{"k smaller or equal than nums", []int{2, 2, 3, 2, 3}, 2, 0},
}

func TestMaxOperations(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MaxOperations(tc.nums, tc.k)
			if got != tc.expect {
				t.Fatalf("MaxOperations(%v, %v) = %v expected: %v", tc.nums, tc.k, got, tc.expect)
			}
		})
	}
}

func TestMaxOperationsWithHashMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MaxOperationsWithHashMap(tc.nums, tc.k)
			if got != tc.expect {
				t.Fatalf("MaxOperationsWithHashMap(%v, %v) = %v expected: %v", tc.nums, tc.k, got, tc.expect)
			}
		})
	}
}
