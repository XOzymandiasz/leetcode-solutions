package _000013_romanToInteger

import "strings"

func romanToInt(s string) int {
	numerals := []string{"CM", "CD", "XC", "XL", "IX", "IV", "M", "D", "C", "L", "X", "V", "I"}
	values := []int{900, 400, 90, 40, 9, 4, 1000, 500, 100, 50, 10, 5, 1}

	result := 0

	for len(s) > 0 {
		for i := 0; i < len(numerals); i++ {
			if strings.HasPrefix(s, numerals[i]) {
				result += values[i]
				s = strings.TrimPrefix(s, numerals[i])
				break
			}
		}
	}

	return result
}

func slowerRomanToInt(s string) int {
	result := 0
	n := len(s)

	if n == 1 {
		return mapRomanToInt(s[0])
	}

	current := 0
	next := 0

	for i := 0; i < n-1; i++ {
		current = mapRomanToInt(s[i])
		next = mapRomanToInt(s[i+1])
		if current >= next {
			result += current
		} else {
			result -= current
		}
	}
	result += next
	return result
}

func mapRomanToInt(r uint8) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
