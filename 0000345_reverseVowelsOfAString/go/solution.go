package reverseVowelsOfAString

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'u': {},
	'o': {},
	'i': {},
	'A': {},
	'U': {},
	'O': {},
	'I': {},
	'E': {},
}

func isVowel(r byte) bool {
	_, ok := vowels[r]
	return ok
}

func ReverseVowels(s string) string {
	n := len(s)
	reverseVowelString := []byte(s)
	left := 0
	right := n - 1

	for left <= right {
		if !isVowel(reverseVowelString[left]) {
			left++
			continue
		}
		if !isVowel(reverseVowelString[right]) {
			right--
			continue
		}
		reverseVowelString[left], reverseVowelString[right] = reverseVowelString[right], reverseVowelString[left]
		left++
		right--
	}

	return string(reverseVowelString)
}
