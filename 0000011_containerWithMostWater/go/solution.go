package solution

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0

	for left < right {
		x := right - left
		y := min(height[left], height[right])
		P := x * y
		result = max(result, P)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return result
}
