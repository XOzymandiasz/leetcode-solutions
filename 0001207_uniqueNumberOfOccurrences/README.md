# 1207. Unique Number Of Occurrences
[All Solutions](../0001207_uniqueNumberOfOccurrences)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/unique-number-of-occurrences

---

## 🧩 Problem
Given an array of integers `arr`, return `true` if the number of occurrences for each value in the array is unique, or `false` otherwise.

Constraints:
- 1 <= arr.length <= 1000
- -1000 <= arr[i] <= 1000
---

## 🔍 Approach

The solution first creates a frequency map to count the occurrences of each integer. Then, it iterates through these counts to verify that each frequency value appears only once by using a second map

---

## ⏱ Complexity

- **Time:** `O(n)` – iterate single time through `arr` and through `counts`
- **Space:** `O(n)` - space for maps

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime      | Memory           |
|----------|--------------|------------------|
| Go       | ~0 ms (100%) | ~4.5  MB (2.56%) |
---
## 💻 Implementations

## 🐹 Go
```go
func UniqueOccurrences(arr []int) bool {    
    n := len(arr)
    counts := make(map[int]int, n)
	
    for _, num := range arr {
    	counts[num] += 1
    }

    occurred := make(map[int]bool, n)
    for _, count := range counts {
    	if occurred[count] {
    		return false
    	}
    	occurred[count] = true
    }
    return true
}
```
