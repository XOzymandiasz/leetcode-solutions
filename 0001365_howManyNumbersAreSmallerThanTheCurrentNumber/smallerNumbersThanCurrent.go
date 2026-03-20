package _001365_howManyNumbersAreSmallerThanTheCurrentNumber

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)

	m := make(map[int]int)

	for i, v := range sorted {
		if _, exists := m[v]; !exists {
			m[v] = i
		}
	}

	answer := make([]int, n)
	for i, v := range nums {
		answer[i] = m[v]
	}

	return answer
}
