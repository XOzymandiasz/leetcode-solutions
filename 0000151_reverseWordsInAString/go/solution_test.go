package solution

import (
	"strings"
	"testing"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	n := len(s)
	answer := []byte(s)
	i := 0
	jEnd := n
	for i < n {
		lastSpace := jEnd
		jStart := jEnd - 1
		for s[jStart] != ' ' {
			jStart -= 1
		}
		lastSpace = jStart + 1
		for jStart < jEnd {
			answer[i] = s[jStart]
			jStart++
			i++
		}
		if i == n {
			break
		}
		answer[i] = ' '
		i++
		jEnd = lastSpace
	}

	return string(answer)
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		//                 isie      js  je
		{"base", "the sky is blue", "blue is sky the"},
		{"base2", "  hello world  ", "world hello"},
		{"base3", "a good   example", "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Fatalf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
