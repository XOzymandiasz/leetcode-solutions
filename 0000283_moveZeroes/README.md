# 283. Move Zeroes
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/move-zeroes

---
## 🧩 Problem
You are given an integer array `nums`.

Modify the array in-place so that:
- all non-zero values appear at the beginning, in the same order as in the input,
- all zero values are moved to the end.

The operation must not use additional memory.
Constraints:
- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

---

## 🔍 Approach
Use two indices:

- one iterates over the array,
- the other tracks the position where the next non-zero value should be written.

As the array is scanned, each non-zero element is copied forward.
Once all values are processed, the remaining part of the array is filled with zeroes.
---

## ⏱ Complexity

- **Time:** `O(n)` - single pass through array `nums`
- **Space:** `O(1)` - no extra space

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~8.3  MB (90.03%) |

## 💻 Implementations

### 🐹 Go
```go
func MoveZeroes(nums []int) {
    n := len(nums)
    j := 0
    i := 0
    for j < n {
        if i < n {
            actual := nums[i]
            if actual != 0 {
                nums[j] = actual
                j++
            }
            i++
        } else {
            nums[j] = 0
            j++
        }
    }
}
```