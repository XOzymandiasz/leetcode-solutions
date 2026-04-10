package solution

func FindMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	actualSum := 0

	for i := 0; i < k; i++ {
		actualSum += nums[i]
	}
	maxSum = actualSum
	for j := 1; j <= len(nums)-k; j++ {
		actualSum = actualSum - nums[j-1] + nums[j+k-1]
		maxSum = max(maxSum, actualSum)
	}

	return float64(maxSum) / float64(k)
}
