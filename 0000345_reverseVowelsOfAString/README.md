# 345. Reverse Vowels of a String
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/reverse-vowels-of-a-string

---

## 🧩 Problem
Given string `s`, return string with reversed vowels.

Constraints:
- 1 <= s.length <= 3*10^5.
- s consist of ASCI characters.

---

## 🔍 Approach
Two pointers traverse the array from the left and right, searching for vowels. 
When both point to vowels, they are swapped. 

---

## ⏱ Complexity

- **Time:** `O(n)` – two pointers traverse the array linearly
- **Space:** `O(n)` - additional memory for result

---
## ⚡ Performance

| Language    | Runtime        | Memory           |
|-------------|----------------|------------------|
| Go          | ~0 ms (100%)   | ~6.1  MB (54.9%) |
---
## 💻 Implementations

## 🐹 Go
```go
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
```
