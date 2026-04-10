package solution

func IsPalindrome(x int) bool {
	original := x
	reversed := 0

	if original < 0 {
		return false
	}

	for x > 0 {
		reversed *= 10
		reversed += x % 10
		x /= 10
	}
	return original == reversed
}
