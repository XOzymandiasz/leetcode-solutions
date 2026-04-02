package greatestCommonDivisors_test

import (
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001071_greatestCommonDivisorOfStrings/go"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{name: "basic example", str1: "ABCABC", str2: "ABC", want: "ABC"},
		{name: "odd number of letters", str1: "ABABAB", str2: "ABAB", want: "AB"},
		{name: "same strings", str1: "ZUBR", str2: "ZUBR", want: "ZUBR"},
		{name: "one letter not odd", str1: "AAAAA", str2: "AA", want: "A"},
		{name: "no divisor", str1: "ZUBR", str2: "RBUZ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.GcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Fatalf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
