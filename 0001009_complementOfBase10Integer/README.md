# 1009. Complement of Base 10 Integer

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/complement-of-base-10-integer

---

## 🧩 Problem
Given an integer `n`, return its **bitwise complement**.

The complement of an integer is obtained by flipping all bits in its binary representation **without leading zeros**.

Constraints:
- `0 <= n < 10^9`

---

## 🔍 Approach

Use bitwise operations to process the number one bit at a time.

Check each bit from least significant to most significant.
If the current bit is 0, set that bit in the result.
If the current bit is 1, clear it from n and continue.

Track the current bit position during the process and return the constructed complement.

---

## ⏱ Complexity

- **Time:** `O(log n)` – number of bits in `n`
- **Space:** `O(1)` – constant extra space

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment.

| Language | Runtime        | Memory           |
|----------|----------------|------------------|
| Go       | ~0 ms (100%)   | ~3.8 MB (99.29%) |

---

## 💻 Implementations

### 🐹 Go
```go
func BitwiseComplement(n int) int {
    if n == 0 {
    	return 1
    }
    answer := 0
    i := 0
    for n != 0 {
    	if n&(1<<i) == 0 {
            answer |= 1 << i
    	} else {
    	    n &^= 1 << i
    	}
    	i++
    }
    return answer
}