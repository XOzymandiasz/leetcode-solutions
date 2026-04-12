package solution

import "sort"

func MaxOperations(nums []int, k int) int {
	left, right := 0, len(nums)-1
	result := 0
	sort.Ints(nums)
	for right > 0 && nums[right] >= k {
		right--
	}
	for left < right {
		if nums[left]+nums[right] == k {
			result++
			left++
			right--
			continue
		}
		if nums[right] > k-nums[left] {
			right--
		} else {
			left++
		}
	}

	return result
}

func MaxOperationsWithHashMap(nums []int, k int) int {
	occurrences := make(map[int]int, min(len(nums), k))
	result := 0

	for _, num := range nums {
		if num >= k {
			continue
		}
		if occurrences[k-num] > 0 {
			occurrences[k-num]--
			result++
			continue
		}
		occurrences[num]++
	}

	return result
}
