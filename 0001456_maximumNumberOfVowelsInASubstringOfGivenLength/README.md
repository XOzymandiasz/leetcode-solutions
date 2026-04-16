# 1456. Maximum Number of Vowels in a Substring of Given Length
[All Solutions](../)  

🟡 Difficulty: Medium

🔗 https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length

---
## 🧩 Problem
Given a string `s` and an integer `k`, return maximum number vowels in range k.

Constraints:
- `1 <= s.length <= 10^5`
- `1 <= k <= s.length`
- `s consist only lowercase English letters`

---

## 🔍 Approach
Use a sliding window of size `k` to track the number of vowels in the current substring.  
Initialize the count for the first window, then slide the window one character at a time by adding the next character and removing the previous one.  
Keep track of the maximum number of vowels seen during the process.
---

## ⏱ Complexity

- **Time:** `O(n)` - single pass through string
- **Space:** `O(1)` - only fixed-size map

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime         | Memory           |
|----------|-----------------|------------------|
| Go       | ~20 ms (26.45%) | ~6.7  MB (62.4%) |

## 💻 Implementations

### 🐹 Go
```go
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
```