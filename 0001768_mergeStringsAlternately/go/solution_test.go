package mergeStringAlternately_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001768_mergeStringsAlternately/go"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{"base", "abc", "def", "adbecf"},
		{"zubr", "zubr", "zubr", "zzuubbrr"},
		{"one word longer", "abcdef", "def", "adbecfdef"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := solution.MergeAlternately(tc.word1, tc.word2)

			if got != tc.want {
				t.Errorf("given: word1 = %v, word2 = %v got: %v; want: %v", tc.word1, tc.word2, got, tc.want)
			}
		})
	}
}
