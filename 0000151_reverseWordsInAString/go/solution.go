package solution

import "strings"

func ReverseWords(s string) string {
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
