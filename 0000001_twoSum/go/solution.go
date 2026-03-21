package twosum

// TwoSum returns indices of two numbers whose sum equals the target.
// Uses a map to store visited numbers and find complements in O(1),
// resulting in an overall O(n) time complexity and O(n) space complexity.
func TwoSum(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))

	for i, n := range nums {
		if idx, exist := visited[target-n]; exist {
			return []int{idx, i}
		}

		visited[n] = i
	}

	return []int{}
}
