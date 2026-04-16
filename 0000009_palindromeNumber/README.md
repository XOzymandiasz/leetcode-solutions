# 9. Palindrome Number
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/palindrome-number

---
## 🧩 Problem
Given an integer `x`, return `true` if `x` is a palindrome, otherwise `false`.

Constraints:
- `-2^31 <= x <= 2^31-1`
---

## 🔍 Approach
The solution iterates through the digits of `x`, constructing the reversed number.  
Finally, it compares the reversed value with the original.

## ⏱ Complexity

- **Time:** `O(log(x))` – one iteration per digit
- **Space:** `O(1)` – only space for reversed digit
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~6.0  MB (96.61%) |

## 💻 Implementations

### 🐹 Go
```go
func IsPalindrome(x int) bool {
    original := x
    reversed := 0

    if original < 0 {
        return false
    }

    for x > 0 {
        reversed *= 10
        reversed += x % 10
        x /= 10
    }
    return original == reversed
}
```