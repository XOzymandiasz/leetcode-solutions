# 13. Roman to Integer
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/roman-to-integer

---
## 🧩 Problem
Given Roman number as string `s`, return Arabic number as integer
Constraints:
- `1 <= s.length <= 15`
- `s` contains only the characters (`I`, `V`, `X`, `L`, `C`, `D`, `M`).
- `s` is in range `[1,3999]`
---

## 🔍 Approach
The solution uses a greedy strategy.

We prepare two arrays:
- `numerals` – Roman numeral patterns (including subtractive cases like `IV`, `IX`, etc.)
- `values` – corresponding integer values

While the string `s` is not empty:
1. Iterate through all numeral patterns.
2. Check if `s` starts with the current numeral.
3. If it does:
    - Add its value to the result.
    - Remove the matched prefix from `s`.
4. Repeat until the string is fully processed.

The numerals are ordered from largest to smallest, ensuring correct matching.

## ⏱ Complexity

- **Time:** `O(n)` – each character is processed once
- **Space:** `O(1)` – only fixed-size arrays are used
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~4.7  MB (92.54%) |

## 💻 Implementations

### 🐹 Go
```go
func RomanToInt(s string) int {
    numerals := []string{"CM", "CD", "XC", "XL", "IX", "IV", "M", "D", "C", "L", "X", "V", "I"}
    values := []int{900, 400, 90, 40, 9, 4, 1000, 500, 100, 50, 10, 5, 1}

    result := 0
    for len(s) > 0 {
        for i := 0; i < len(numerals); i++ {
            if strings.HasPrefix(s, numerals[i]) {
                result += values[i]
                s = strings.TrimPrefix(s, numerals[i])
                break
            }
        }
    }
    return result
}
```