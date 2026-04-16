# 11. Container With Most Water
[All Solutions](../)

🟠 Difficulty: Medium

🔗 https://leetcode.com/problems/container-with-most-water

---
## 🧩 Problem
Given an array of integer `height`, return maximum possible area.
- `height[i]` represents `y`
- `i` represents `x`

Constraints:
- `2 <= heigth.length <= 10^5`
- `0 <= height[i] <= 10^4`
---

## 🔍 Approach
Use a two-pointer technique: start with one pointer at the beginning and one at the end.
At each step, calculate the area using the shorter height and update the maximum.
Move the pointer pointing to the smaller height, since increasing height is the only way to potentially get a larger area.

Repeat until the pointers meet.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through array
- **Space:** `O(1)` – only space for reversed digit
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~9.3  MB (92.17%) |

## 💻 Implementations

### 🐹 Go
```go
func MaxArea(height []int) int {
    left := 0
    right := len(height) - 1
    result := 0

    for left < right {
        x := right - left
        y := min(height[left], height[right])
        P := x * y
        result = max(result, P)
        if height[left] < height[right] {
            left += 1
        } else {
            right -= 1
        }
    }
    return result
}

```