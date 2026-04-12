# 495. Teemo Attacking

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/teemo-attacking

---
## 🧩 Problem
Given an array timeSeries representing attack times and an integer duration, return the total time the target is poisoned.

Constraints:
- `1 <= timeSeries.length <= 10^4`
- `0 <= timeSeries[i], duration <= 10^7`
---

## 🔍 Approach
Solution iterate through the array and calculate how much each attack contributes to the total poisoned time. 
If the gap between consecutive attacks is greater than or equal to duration, we add the full duration; 
otherwise, we only add the gap to avoid double-counting overlapping poison time. 
After processing all pairs, we add the full duration for the last attack. 
This works because overlapping intervals should only be counted once.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through the array
- **Space:** `O(1)` – constant extra space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~8.6  MB (75.34%) |

## 💻 Implementations

### 🐹 Go
```go
func FindPoisonedDuration(timeSeries []int, duration int) int {
    totalDuration := 0
    n := len(timeSeries)

    for i := 0; i < n-1; i++ {
        interval := timeSeries[i+1] - timeSeries[i]
        if (interval) > duration {
            totalDuration += duration
        } else {
            totalDuration += interval
        }
    }

    return totalDuration + duration
}

```