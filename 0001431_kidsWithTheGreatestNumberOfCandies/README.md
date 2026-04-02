# 1431. Kids With the Greatest Number of Candies

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/kids-with-the-greatest-number-of-candies

---
## 🧩 Problem
Given an array of integers `candies`, where `candies[i]` represents the number of candies the i-th child has,
and an integer `extraCandies`, return a boolean array where each element indicates whether the i-th child can have
the greatest number of candies after receiving all the extra candies.

- Multiple children can have the greatest number of candies

Constraints:
- `2 <= candies.length <= 100`
- `1 <= candies[i] <= 100`
- `1 <= extraCandies <= 50`

---

## 🔍 Approach
First, find the maximum number of candies and subtract `extraCandies` from it.  
Then, check if each child already has at least that amount.
---

## ⏱ Complexity

- **Time:** `O(n)` - one pass through array to search max value and second to build answer
- **Space:** `O(n)` - the result contains answers for every child

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~4.1  MB (91.44%) |

## 💻 Implementations

### 🐹 Go
```go
func KidsWithCandies(candies []int, extraCandies int) []bool {
    n := len(candies)
    var answer []bool
    maximum := maxValue(candies) - extraCandies
    for i := 0; i < n; i++ {
        answer = append(answer, candies[i] >= maximum)
    }
    return answer
}    
    
func maxValue(a []int) int {   
    maximum := a[0]   
    for i := 1; i < len(a); i++ {
        if a[i] > maximum {
            maximum = a[i]
        }   		
    }
    return maximum
}
```