package _001768_mergeStringsAlternately

func MergeAlternately(word1 string, word2 string) string { //2ms runtime 4.0MB memory
	n1 := len(word1)
	n2 := len(word2)
	n := max(n1, n2)

	answer := make([]byte, n1+n2)
	j := 0
	for i := 0; i < n; i++ {
		if i < n1 {
			answer[j] = word1[i]
			j++
		}
		if i < n2 {
			answer[j] = word2[i]
			j++
		}
	}

	return string(answer)
}

func notMineSolution(word1 string, word2 string) string { //0ms runtime 4.8MB memory
	var result string
	n := len(word1) + len(word2)
	mn := min(len(word1), len(word2))
	for i, j, k := 0, 0, 0; i < n; i++ {
		if k == len(word2) {
			break
		}
		if i%2 == 0 {
			if j == len(word1) {
				break
			}
			result += string(word1[j])
			j += 1
		} else {
			if k == len(word2) {
				break
			}
			result += string(word2[k])
			k += 1
		}

	}
	if len(word1) > len(word2) {
		result += word1[mn:]
	} else if len(word1) < len(word2) {
		result += word2[mn:]
	}
	return result
}
