package solution_test

import (
	"reflect"
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000700_searchInABinarySearchTree/go"
)

func insert(root *solution.TreeNode, val int) *solution.TreeNode {
	if root == nil {
		return &solution.TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

func buildBST(nums []int) *solution.TreeNode {
	var root *solution.TreeNode
	for _, num := range nums {
		root = insert(root, num)
	}
	return root
}

func toPreorderSlice(root *solution.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}
	result = append(result, toPreorderSlice(root.Left)...)
	result = append(result, toPreorderSlice(root.Right)...)

	return result
}

func TestSearchBST(t *testing.T) {
	tests := []struct {
		name string
		root []int
		val  int
		want []int
	}{
		{name: "found root{search 4}", root: []int{4, 2, 7, 1, 3}, val: 4, want: []int{4, 2, 1, 3, 7}},
		{name: "found left subtree{search 2}", root: []int{4, 2, 7, 1, 3}, val: 2, want: []int{2, 1, 3}},
		{name: "found right node{search 7}", root: []int{4, 2, 7, 1, 3}, val: 7, want: []int{7}},
		{name: "not found{search 5}", root: []int{4, 2, 7, 1, 3}, val: 5, want: nil},
		{name: "empty tree{search 1}", root: []int{}, val: 1, want: nil},
		{name: "single node found{search 1}", root: []int{1}, val: 1, want: []int{1}},
		{name: "single node not found{search 2}", root: []int{1}, val: 2, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := toPreorderSlice(solution.SearchBST(buildBST(tc.root), tc.val))

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SearchBST(%v, %d) = %v, want %v", tc.root, tc.val, got, tc.want)
			}
		})
	}
}
