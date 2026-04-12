package addTwoNumbers_test

import (
	"reflect"
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000002_addTwoNumbers/go"
)

func buildList(nums []int) *solution.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &solution.ListNode{Val: nums[0], Next: nil}
	current := head

	for _, num := range nums[1:] {
		current.Next = &solution.ListNode{Val: num, Next: nil}
		current = current.Next
	}

	return head
}

func toSlice(node *solution.ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{"base{321 + 465 = 786}", []int{1, 2, 3}, []int{5, 6, 4}, []int{6, 8, 7}},
		{"carry over{999 + 999 = 1998}", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},
		{"carry over with longer first list{1111 + 999 = 2110}", []int{1, 1, 1, 1}, []int{9, 9, 9}, []int{0, 1, 1, 2}},
		{"empty{0+0=0}", []int{}, []int{}, []int{0}},
		{"empty second list{321 + 0 = 321}", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"empty first list{0 + 321 = 321}", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := toSlice(solution.AddTwoNumbers(buildList(tc.l1), buildList(tc.l2)))

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("AddTwoNumber(%v, %v) = %v, want %v", tc.l1, tc.l2, got, tc.want)
			}
		})
	}
}
