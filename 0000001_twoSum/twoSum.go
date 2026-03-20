package _000001_twoSum

func TwoSum(nums []int, target int) []int {
	potential := make(map[int]int)

	for i, n := range nums {
		if val, exist := potential[target-n]; exist {
			return []int{val, i}
		}

		potential[n] = i
	}

	return []int{1, -1}
}
