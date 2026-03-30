package mergeStringAlternately

import "strings"

func MergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	n := max(n1, n2)

	var answer strings.Builder
	answer.Grow(n1 + n2)
	for i := 0; i < n; i++ {
		if i < n1 {
			answer.WriteByte(word1[i])
		}
		if i < n2 {
			answer.WriteByte(word2[i])
		}
	}

	return answer.String()
}
