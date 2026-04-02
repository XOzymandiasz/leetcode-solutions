package solution

func KidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	var answer []bool
	maxCandies := maxValue(candies) - extraCandies
	for i := 0; i < n; i++ {
		answer = append(answer, candies[i] >= maxCandies)
	}
	return answer
}

func maxValue(a []int) int {
	maximum := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > maximum {
			maximum = a[i]
		}
	}
	return maximum
}
