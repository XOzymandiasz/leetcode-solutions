package solution

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'u': {},
	'o': {},
	'i': {},
}

func isVowel(r byte) bool {
	_, ok := vowels[r]
	return ok
}

func MaxVowels(s string, k int) int {
	maxNumberOfVowels, currentNumberOfVowels := 0, 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currentNumberOfVowels++
		}
	}
	maxNumberOfVowels = currentNumberOfVowels

	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			currentNumberOfVowels++
		}
		if isVowel(s[i-k]) {
			currentNumberOfVowels--
		}
		maxNumberOfVowels = max(currentNumberOfVowels, maxNumberOfVowels)
	}

	return maxNumberOfVowels
}
