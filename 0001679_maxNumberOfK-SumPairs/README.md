# 1679. Max Number Of K-Sum Pairs
[All Solutions](../)  

🟡 Difficulty: Medium

🔗 https://leetcode.com/problems/max-number-of-k-sum-pairs

---
## 🧩 Problem
Given an array of integers nums and an integer `k`, return number of operations you can perform to delete all k-sum pairs.

Constraints:
- `1 <= nums.length <= 10^5`
- `1 <= nums[i], k <= 10^9`

---

## 🔍 Approach
Sort the array and use a two-pointer technique.
Place one pointer at the beginning and one at the end.
- If `nums[left] + nums[right] == k` then form a pair and move both pointers
- If the sum is too small → move left forward
- If the sum is too large → move right backward

Repeat until pointers meet, counting valid pairs.
---

## ⏱ Complexity

- **Time:** `O(nlogn)` - time complexity of introsort
- **Space:** `O(logn)` - space for introsort

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime         | Memory             |
|----------|-----------------|--------------------|
| Go       | ~77 ms (80.29%) | ~9.42  MB (85.27%) |

## 💻 Implementations

### 🐹 Go
```go
func MaxOperations(nums []int, k int) int {
    left, right := 0, len(nums)-1
    result := 0
    sort.Ints(nums)
    for right > 0 && nums[right] >= k {
        right--
    }
    for left < right {
        if nums[left]+nums[right] == k {
            result++
            left++
            right--
            continue
        }
        if nums[right] > k-nums[left] {
            right--
        } else {
            left++
        }
    }
    return result
}
```