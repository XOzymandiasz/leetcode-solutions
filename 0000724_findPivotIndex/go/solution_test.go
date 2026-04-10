package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000724_findPivotIndex/go"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "first", nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "second", nums: []int{1, 2, 3}, want: -1},
		{name: "third", nums: []int{2, 1, -1}, want: 0},
		{name: "pivot_in_middle", nums: []int{1, 2, 1}, want: 1},
		{name: "single_element", nums: []int{10}, want: 0},
		{name: "empty_slice", nums: []int{}, want: -1},
		{name: "pivot_at_last_index", nums: []int{1, -1, 0}, want: 2},
		{name: "all_zeros", nums: []int{0, 0, 0, 0}, want: 0},
		{name: "negative_numbers", nums: []int{-1, -1, -1, -1, 0, 1}, want: 1},
		{name: "no_pivot_with_negatives", nums: []int{-1, -1, 0, 1, 1, 0}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution.PivotIndex(tt.nums)
			if got != tt.want {
				t.Fatalf("pivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
