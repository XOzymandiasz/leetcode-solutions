package solution_test

import (
	"reflect"
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000283_moveZeroes/go"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{"base", []int{3, 0, 0, 4}, []int{3, 4, 0, 0}},
		{"base 2", []int{0, 2, 1, 0, 68, 3, 0, 0, 4}, []int{2, 1, 68, 3, 4, 0, 0, 0, 0}},
		{"single zero", []int{0}, []int{0}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var given []int
			copy(given, tc.nums)
			solution.MoveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Errorf("given: nums = %v, expected: %v, got: %v", given, tc.want, tc.nums)
			}
		})
	}
}
