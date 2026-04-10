# 1732. Find The Highest Altitude

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/find-the-highest-altitude

---
## 🧩 Problem
Given an integer array gain, where gain[i] is the net gain in altitude between points i and i + 1, return the highest altitude reached.
- Start from 0

Constraints:
- `1 <= gain.length <= 100`
- `-100 <= gain[i] <= 100`
---

## 🔍 Approach
Track the current altitude while iterating through gain.
After each step, update the highest altitude seen so far.
Since the biker starts at altitude `0`, initialize both current and highest altitude with `0`.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through array
- **Space:** `O(1)` – only constant space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~4.0  MB (45.99%) |

## 💻 Implementations

### 🐹 Go
```go
func LargestAltitude(gain []int) int {
    altitude := 0
    highest := 0

    for _, point := range gain {
        altitude = altitude + point
        highest = max(highest, altitude)
    }

    return highest
}
```