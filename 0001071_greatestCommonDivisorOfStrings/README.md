# 1071. Greatest Common Divisor Of Strings

🟢 Difficulty: Easy

---
## 🧩 Problem
Given two string `s` and `t`, return largest string that both `s` and `t` can be formed by concatenating returned string.

Constraints
- `1 <= s.length, t.length <= 1000`
- `s` and `t` consist of English uppercase letters.

## 🔍 Approach
The solution checks `s+t` is equal to `t+s` and then use Euclides algorithm to find the greatest common divisor. 

## ⏱ Complexity

- **Time:** `O(n + m)` - comparison of two concatenated strings
- **Space:** `O(n + m)` – due to temporary strings created by `s + t` and `t + s`
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~4.3  MB (76.46%) |

## 💻 Implementations

### 🐹 Go
```go
func GcdOfStrings(str1 string, str2 string) string {
    k1 := len(str1)
    k2 := len(str2)

    if str1+str2 != str2+str1 {
        return ""
    }	

    for k2 != 0 {
        k1, k2 = k2, k1%k2
    }
    return str1[:k1]
}
```