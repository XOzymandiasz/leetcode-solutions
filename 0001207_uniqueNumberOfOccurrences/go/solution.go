package solution

func UniqueOccurrences(arr []int) bool {
	n := len(arr)
	counts := make(map[int]int, n)

	for _, num := range arr {
		counts[num] += 1
	}

	occurred := make(map[int]bool, n)
	for _, count := range counts {
		if occurred[count] {
			return false
		}
		occurred[count] = true
	}

	return true
}
