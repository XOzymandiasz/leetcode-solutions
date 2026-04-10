package solution

const ZERO = '0'
const ONE = '1'

func MinOperations(s string) int {
	counter := 0
	var last rune
	for _, r := range s {
		if r == last {
			counter++
			if r == ZERO {
				last = ONE
			} else {
				last = ZERO
			}
		} else {
			last = r
		}
	}

	return min(counter, len(s)-counter)
}
