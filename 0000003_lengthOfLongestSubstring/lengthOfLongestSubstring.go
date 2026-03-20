package _000003_lengthOfLongestSubstring

import "fmt"

func lengthOfLongestSubstring(s string) int {
	alphabet := make(map[rune]int)
	length := 0
	longest := 0
	runeRepeat := false

	for i, r := range s {
		val, exist := alphabet[r]

		j := 0
		if exist && val != 0 {
			length -= alphabet[r]

			runeRepeat = true
			for {
				if j == alphabet[r]-1 {
					if alphabet[rune(s[j])] >= alphabet[r] {
						alphabet[rune(s[j])] = 0
						//delete(alphabet, rune(s[j]))
					}
					j++
					break
				}
				if alphabet[rune(s[j])] >= alphabet[r] {
					alphabet[rune(s[j])] = 0
					//delete(alphabet, rune(s[j]))
				}
				j++
			}

		}

		length++
		alphabet[r] = i + 1
		if runeRepeat {
			for a := range alphabet {
				alphabet[a] -= j
				if alphabet[a] < 0 {
					alphabet[a] = 0
				}
			}
			runeRepeat = false
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

func printAlphabet(alphabet map[rune]int) {
	for a, i := range alphabet {
		fmt.Print(a)
		fmt.Print(" : ")
		fmt.Println(i)
	}
}
