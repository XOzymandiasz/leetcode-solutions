# 643. Maximum Average Subarray I
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/maximum-average-subarray-i

---
## 🧩 Problem
Given an integer array `nums` and an integer `k`, return the maximum average value of any contiguous subarray of length `k`.

Constraints:
- `1 <= k <= n <= 10^5`
- `-10^4 <= nums[i] 10^4`
---

## 🔍 Approach
Use a sliding window of size `k`.

First, compute the sum of the first `k` elements.  
Then, slide the window across the array by subtracting the element leaving the window and adding the next element.

Track the maximum sum during the process and return the maximum average.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through the array
- **Space:** `O(1)` – constant extra space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~9.8  MB (56.46%) |

## 💻 Implementations

### 🐹 Go
```go
func FindMaxAverage(nums []int, k int) float64 {
    maxSum := 0
    actualSum := 0

    for i := 0; i < k; i++ {
        actualSum += nums[i]
    }
    maxSum = actualSum
    for j := 1; j <= len(nums)-k; j++ {
        actualSum = actualSum - nums[j-1] + nums[j+k-1]
        maxSum = max(maxSum, actualSum)
    }

    return float64(maxSum) / float64(k)
}
```