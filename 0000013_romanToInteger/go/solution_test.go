package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000013_romanToInteger/go"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{"One", "I", 1},
		{"Four", "IV", 4},
		{"Nine", "IX", 9},
		{"Forty", "XL", 40},
		{"Ninety", "XC", 90},
		{"Four Hundred", "CD", 400},
		{"Nine Hundred", "CM", 900},
		{"Two Thousand Twenty Four", "MMXXIV", 2024},
		{"One Thousand Nine Hundred Eighty Seven", "MCMLXXXVII", 1987},
		{"One Thousand Six Hundred Sixty Six", "MDCLXVI", 1666},
		{"Three Thousand Nine Hundred Ninety Nine", "MMMCMXCIX", 3999},
		{"One Thousand Nine Hundred Ninety Four", "MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.RomanToInt(tc.s)
			if got != tc.want {
				t.Errorf("given: s=%v, got: %v; want: %v", tc.s, got, tc.want)
			}
		})
	}
}
