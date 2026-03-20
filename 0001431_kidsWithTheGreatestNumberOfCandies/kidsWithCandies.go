package _001431_kidsWithTheGreatestNumberOfCandies

func KidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	var answer []bool
	maximum := maxValue(candies) - extraCandies
	for i := 0; i < n; i++ {
		answer = append(answer, candies[i] >= maximum)
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
