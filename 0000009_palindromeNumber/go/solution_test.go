package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0000009_palindromeNumber/go"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		want bool
	}{
		{"single digit", 1, true},
		{"zero", 0, true},
		{"palindrome", 12321, true},
		{"second palindrome", 1001, true},
		{"third palindrome", 10101, true},
		{"not palindrome", 12123, false},
		{"negative", -11, false},
		{"ends with zero", 1000, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.IsPalindrome(tc.x)
			if tc.want != got {
				t.Fatalf("given: x = %v expected: %v, got: %v", tc.x, tc.want, got)
			}
		})
	}
}
