# 724. Find Pivot Index
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/find-pivot-index

---
## 🧩 Problem
Given an integer array `nums`, return the pivot index.
- The pivot index is the index where the sum of all elements strictly to the left
is equal to the sum of all elements strictly to the right.
- If no such index exists, return `-1`.

Constraints:
- `1 <= nums.length <= 10^4`
- `-1000 <= nums[i] <= 1000`
---

## 🔍 Approach
First, compute the total sum of the array.

Then iterate through the array while maintaining a running `leftSum`.  
At each index, subtract the current element from the total sum to get `rightSum`.

If `leftSum == rightSum`, return the current index as the pivot.

Otherwise, add the current element to `leftSum` and continue.

## ⏱ Complexity

- **Time:** `O(n)` – one iteration per digit
- **Space:** `O(1)` – constant extra space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~7.7  MB (98.01%) |

## 💻 Implementations

### 🐹 Go
```go
func PivotIndex(nums []int) int {
    n := len(nums)
    rightSum := 0
    leftSum := 0
    pivotIdx := -1

    for i := 0; i < n; i++ {
        rightSum += nums[i]
    }
    for i := 0; i < n; i++ {
        rightSum -= nums[i]
        if leftSum == rightSum {
            pivotIdx = i
            break
        }
        leftSum += nums[i]
    }

    return pivotIdx
}
```