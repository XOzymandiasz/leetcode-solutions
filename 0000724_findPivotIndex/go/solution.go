package solution

func PivotIndex(nums []int) int {
	n := len(nums)
	rightSum := 0
	leftSum := 0
	pivotIdx := -1

	for i := 0; i < n; i++ {
		rightSum += nums[i]
	}
	for i := 0; i < n; i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			pivotIdx = i
			break
		}
		leftSum += nums[i]
	}

	return pivotIdx
}
