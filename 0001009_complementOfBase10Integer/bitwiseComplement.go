package _001009_complementOfBase10Integer

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	answer := 0
	i := 0
	for n != 0 {
		if n&(1<<i) == 0 {
			answer |= 1 << i
		} else {
			n &^= 1 << i
		}
		i++
	}
	return answer
}
