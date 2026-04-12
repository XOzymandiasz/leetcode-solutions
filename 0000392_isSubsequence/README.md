# 392. Is Subsequence

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/is-subsequence

---
## 🧩 Problem
Given two strings s and t, return true if s is a subsequence of t, otherwise return false.

Constraints:
- `0 <= s.length <= 100`
- `0 <= t.length <= 104`
- `s and t consist only of lowercase English letters.`

---

## 🔍 Approach
We use two pointers to iterate through both strings. 
One pointer tracks the current character in s, and the other scans through t. 
Whenever characters match, we move the pointer in s forward. 
If we reach the end of s, it means all characters were found in order, so we return true; 
otherwise, after scanning t, we return false.
---

## ⏱ Complexity

- **Time:** `O(n)` - solution scan string t once
- **Space:** `O(1)` - no extra space

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~3.9  MB (61.16%) |

## 💻 Implementations

### 🐹 Go
```go
func IsSubsequence(s string, t string) bool {
    sLen := len(s)
    tLen := len(t)
    si := 0
    tj := 0
    if (tLen == 0 && sLen == 0) || sLen == 0 {
        return true
	} else if tLen == 0 {
        return false
    }

    for tj < tLen {
        if s[si] == t[tj] {
          si++
            if si == sLen {
                return true
            }
        }
        tj++
    }
    return false
}
```