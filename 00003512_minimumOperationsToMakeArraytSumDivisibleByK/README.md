# 3512. Minimum Operations to Make Array Sum Divisible by K
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/

---

## 🧩 Problem

Given an integer array `nums` and an integer `k`, you can perform the following operation any number of times:

- Select an index `i` and replace `nums[i]` with `nums[i] - 1`.

Return the minimum number of operations required to make the sum of the array divisible by `k`.

Examples:

```text
Input: nums = [3,9,7], k = 5
Output: 4

Input: nums = [4,1,3], k = 4
Output: 0

Input: nums = [3,2], k = 6
Output: 5
```

Constraints:

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 1000`
- `1 <= k <= 100`

---

## 🔍 Approach

Calculate the sum of all numbers and find its remainder after division by `k`.

Each operation decreases the array's sum by exactly `1`. If the remainder is `r`, the closest lower sum divisible by `k` is therefore reached after exactly `r` operations. When the remainder is `0`, the sum is already divisible by `k` and no operations are needed.

---

## ⏱ Complexity

- **Time:** `O(n)` – every element is visited once while calculating the sum.
- **Space:** `O(1)` – no additional space depending on the input size is used.

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory              |
|----------|----------------|---------------------|
| C#       | ~3 ms (43.33%) | ~54.06  MB (12.38%) |

---

## 💻 Implementations

### C#

```csharp
namespace Solution;

public class Solution
{
    public int MinOperations(int[] nums, int k)
    {
        return nums.Sum() % k;
    }
}
```
