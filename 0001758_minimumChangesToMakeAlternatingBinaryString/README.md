# 9. Palindrome Number

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string

---
## 🧩 Problem
Given an string `s`, return number of changes needed to make alternating binary string.

Constraints:
- `1 <= s.length <= 10^4`
- `s` contest only `0` and `1`
---

## 🔍 Approach
Check both alternating patterns (`0101...` and `1010...`).  
Count mismatches for one pattern and derive the other as `n - mismatches`.  
Return the minimum.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through runes
- **Space:** `O(1)` – only constant space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~4.4  MB (32.61%) |

## 💻 Implementations

### 🐹 Go
```go
const ZERO = '0'
const ONE = '1'

func MinOperations(s string) int {
    counter := 0
    var last rune
    for _, r := range s {
        if r == last {
            counter++
            if r == ZERO {
                last = ONE
            } else {
                last = ZERO
            }
        } else {
            last = r
        }
    }
    return min(counter, len(s)-counter)
}
```