package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000643_maximumAverageSubarray/go"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{name: "classic example", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75},
		{name: "all positive", nums: []int{5, 5, 5, 5}, k: 2, want: 5.0},
		{name: "all negative", nums: []int{-1, -2, -3, -4}, k: 2, want: -1.5},
		{name: "single element window", nums: []int{7, 1, 5, 3}, k: 1, want: 7.0},
		{name: "k equals length", nums: []int{2, 2, 2}, k: 3, want: 2.0},
		{name: "increasing sequence", nums: []int{1, 2, 3, 4, 5}, k: 2, want: 4.5},
		{name: "mixed values", nums: []int{0, -1, 2, -3, 4, -5, 6}, k: 3, want: 1.6666667},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution.FindMaxAverage(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findMaxAverage(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
