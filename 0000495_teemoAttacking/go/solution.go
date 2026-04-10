package solution

func FindPoisonedDuration(timeSeries []int, duration int) int {
	totalDuration := 0
	n := len(timeSeries)

	for i := 0; i < n-1; i++ {
		interval := timeSeries[i+1] - timeSeries[i]
		if (interval) > duration {
			totalDuration += duration
		} else {
			totalDuration += interval
		}
	}

	return totalDuration + duration
}
