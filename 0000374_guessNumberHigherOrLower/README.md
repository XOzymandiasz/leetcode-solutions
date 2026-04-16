# 374. Guess Number Higher Or Lower

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/guess-number-higher-or-lower

---

## 🧩 Problem
Given an integer `n` representing the maximum chosen number, return chosen number.

Also given an API `guess(num)` to determine whether the guess is higher, lower, or correct

`guess(num):`
- `-1 if your guess is too high`
- `1  if your guess is too low`
- `0  if equal`


Constraints:
- `1 <= n <= 2^31 - 1`
- `1 <= answer <= n`

---

## 🔍 Approach
Use binary search on range [1, n].
At each step:
- pick middle value
- use `guess(mid)` to determine direction
- narrow search space accordingly

---

## ⏱ Complexity

- **Time:** `O(log(n))` – binary search complexity
- **Space:** `O(1)` - constant memory

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language    | Runtime        | Memory             |
|-------------|----------------|--------------------|
| Go          | ~0 ms (100%)   | ~3.68  MB (99.76%) |
---
## 💻 Implementations

## 🐹 Go
```go
func GuessNumber(n int) int {
    low, high := 1, n
    mid := low + (high-low)/2
    for {
        if guess(mid) == -1 {
            high = mid - 1
        } else if guess(mid) == 1 {
            low = mid + 1
        } else {
            break
        }
        mid = low + (high-low)/2
    }
    
    return mid
}
```
