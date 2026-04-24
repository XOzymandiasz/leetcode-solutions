package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0002390_removingStarsFromAString/go"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{"first case", "xo*zymad**iasz***", "xzymi"},
		{"second case", "bi*z*on*", "bo"},
		{"remove all letters", "xozymadiasz***********", ""},
		{"no stars", "xozymadiasz", "xozymadiasz"},
		{"more stars than letters", "xozymadiasz*************", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.RemoveStars(tc.s)
			if got != tc.expected {
				t.Fatalf("RemoveStars(%q) = %q, expected %q", tc.s, got, tc.expected)
			}
		})
	}
}
