# 1768. Merge Strings Alternately
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/merge-strings-alternately

---
## 🧩 Problem
Given two strings `word1` and `word2`, return a merged string by adding alternating letters.
- Start with `word1` 
- If one string is longer append the remaining letters to the end.

Constraints:
- `1 <= word1.length, word2.length <= 100`
- Words consist of lowercase English letters

---

## 🔍 Approach
The solution iterates from `0` to `max(n,m)`, adding letters alternately to the result.

### Notes
#### 🐹 Go
- The solution uses strings.Builder with Grow to preallocate memory and avoid unnecessary reallocations and copying.

---

## ⏱ Complexity

- **Time:** `O(n+m)` - single pass through both strings
- **Space:** `O(n+m)` - the result contains all characters from both strings

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~3.9  MB (82.61%) |

## 💻 Implementations

### 🐹 Go
```go
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
```