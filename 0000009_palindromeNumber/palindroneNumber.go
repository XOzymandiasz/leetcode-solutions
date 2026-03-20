package _000009_palindromeNumber

func isPalindrome(x int) bool {
	y := 0

	if x < 0 {
		return false
	}
	if x/10 == 0 {
		return true
	}
	if x/100 == 0 {
		return x/10 == x%10
	}
	if x/1000 == 0 {
		return x/100 == x%10
	}

	for x > 0 {
		lastDigit := x % 10
		y += lastDigit
		if y >= x && x == y {
			return true
		}
		x = x / 10

		if x == y {
			return true
		} else if y == 0 {
			return false
		}

		y *= 10
	}

	return false
}
