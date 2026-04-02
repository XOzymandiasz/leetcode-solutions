package solution

func MoveZeroes(nums []int) {
	n := len(nums)
	j := 0
	i := 0
	for j < n {
		if i < n {
			actual := nums[i]
			if actual != 0 {
				nums[j] = actual
				j++
			}
			i++
		} else {
			nums[j] = 0
			j++
		}
	}
}
