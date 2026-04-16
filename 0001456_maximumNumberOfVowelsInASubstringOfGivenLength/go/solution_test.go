package solution_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001456_maximumNumberOfVowelsInASubstringOfGivenLength/go"
)

func TestMaxVowels(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"first case", "abciiueidef", 3, 3},
		{"whole string", "abc", 3, 1},
		{"whole string with last rune as vowel", "cba", 3, 1},
		{"no vowels", "rhythms", 4, 0},
		{"all vowels", "aeiou", 2, 2},
		{"k equals 1 vowel", "abc", 1, 1},
		{"k equals 1 no vowel at first", "zab", 1, 1},
		{"maximum at the end", "xxxaei", 3, 3},
		{"maximum at the beginning", "aeixxx", 3, 3},
		{"sliding window typical", "weallloveyou", 7, 4},
		{"repeated pattern", "abababab", 3, 2},
		{"single char vowel", "a", 1, 1},
		{"single char consonant", "b", 1, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MaxVowels(tc.s, tc.k)
			if got != tc.expected {
				t.Fatalf("MaxVowels(%q, %d) = %d, want %d", tc.s, tc.k, got, tc.expected)
			}
		})
	}
}
