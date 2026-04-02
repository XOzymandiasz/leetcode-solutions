package solution

import "strings"

func RomanToInt(s string) int {
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
