package reverseVowelsOfAString_test

import (
	"testing"

	reverseVowelsOfAString "github.com/XOzymandiasz/leetcode-solutions/0000345_reverseVowelsOfAString/go"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "basic example", in: "helloworld", want: "hollowerld"},
		{name: "same letters", in: "greed", want: "greed"},
		{name: "no vowels", in: "rhythm", want: "rhythm"},
		{name: "only vowels", in: "aeiou", want: "uoiea"},
		{name: "single char vowel", in: "a", want: "a"},
		{name: "single char consonant", in: "b", want: "b"},
		{name: "empty string", in: "", want: ""},
		{name: "mixed case", in: "hEllO", want: "hOllE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowelsOfAString.ReverseVowels(tt.in)
			if got != tt.want {
				t.Fatalf("ReverseVowels(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
