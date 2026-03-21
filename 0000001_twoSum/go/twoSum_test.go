package twosum_test

import (
	"testing"

	twosum "github.com/XOzymandiasz/leetcode-solutions/0000001_twoSum"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{"basic", []int{2, 7, 11, 15}, 9},
		{"unordered", []int{3, 2, 4}, 6},
		{"duplicates", []int{3, 3}, 6},
		{"negative", []int{7, 3, 4, -1}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := twosum.TwoSum(tc.nums, tc.target)

			if len(result) != 2 {
				t.Fatalf("expected 2 indices, got %v", result)
			}

			i, j := result[0], result[1]

			if i == j {
				t.Fatalf("indices must be different: %v", result)
			}

			if i < 0 || i >= len(tc.nums) || j < 0 || j >= len(tc.nums) {
				t.Fatalf("indices out of range: %v", result)
			}

			sum := tc.nums[i] + tc.nums[j]
			if sum != tc.target {
				t.Errorf(
					"nums[%d]=%d + nums[%d]=%d = %d, expected %d",
					i, tc.nums[i],
					j, tc.nums[j],
					sum,
					tc.target,
				)
			}
		})
	}
}
