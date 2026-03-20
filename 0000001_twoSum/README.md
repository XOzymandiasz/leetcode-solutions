# 1. Two Sum

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/two-sum/

---

## 🧩 Problem
Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they add up to `target`.

Constraints:
- Each input has exactly one solution.
- The answer can be returned in any order.

---

## 🔍 Approach
The solution uses a **map** to track previously visited numbers and their indices.

### Key ideas

- For each element `n`, compute its complement: `target - n`
- Check if the complement has already been seen
- If yes — return both indices immediately
- Otherwise — store the current number in the map

### Why it works

Instead of checking all pairs (which would take `O(n²)`), the hash map allows us to:
- store values we've seen
- check complements in constant time `O(1)`

This reduces the overall complexity to linear time.
---

## ⏱ Complexity
- **Time:** `O(n)` - single pass through the array
- **Space:** `O(n)` - additional memory for the hash map

---

## 💻 Implementations

### 🟦 Go
```go
package twosum

func TwoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for i, n := range nums {
		if idx, exist := visited[target-n]; exist {
			return []int{idx, i}
		}

		visited[n] = i
	}

	return []int{}
}
```