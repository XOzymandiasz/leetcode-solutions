package solution

func guess(n int) int {
	if n > 2 {
		return -1
	} else if n == 2 {
		return 0
	}
	return 1
}

func GuessNumber(n int) int {
	low, high := 1, n
	mid := low + (high-low)/2
	for {
		if guess(mid) == -1 {
			high = mid - 1
		} else if guess(mid) == 1 {
			low = mid + 1
		} else {
			break
		}
		mid = (high - low) / 2
	}

	return mid
}
