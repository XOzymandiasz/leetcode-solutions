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

			i := result[0]
			j := result[1]

			if i == j {
				t.Fatalf("indices must be different: %v", result)
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

func TestTwoSum_NoSolution(t *testing.T) {
	result := twosum.TwoSum([]int{3}, 100)

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}
