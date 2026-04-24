# 2390. Removing Stars From A String
[All Solutions](../)  

🟡 Difficulty: Medium

🔗 https://leetcode.com/problems/removing-stars-from-a-string

---

## 🧩 Problem
Given a string `s` consisting of lowercase English letters and `'*'` characters:
- Each `'*'` removes the closest non-star character to its left, as well as the `'*'` itself.
- Return the final string after all removals.

Constraints:
- `1 <= s.length <= 10^5`
- `s` consists of lowercase English letters and `'*'`
- The input is generated such that the operation is always possible (no invalid removals)

---

## 🔍 Approach
The solution uses a **stack-like structure** (slice in Go):

- Iterate through the string from left to right
- If the current character is not `'*'`, push it onto the stack
- If it is `'*'`, remove the last character from the stack (pop)
- Finally, convert the stack to a string

This works because each `'*'` always affects the **most recent valid character**, which matches the **LIFO (Last In, First Out)** behavior of a stack.

## ⏱ Complexity

- **Time:** `O(n)` – single pass through string
- **Space:** `O(n)` - extra space for answer

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory             |
|----------|----------------|--------------------|
| Go       | ~6 ms (71.55%) | ~8.75  MB (84.48%) |
---
## 💻 Implementations

## 🐹 Go
```go
func RemoveStars(s string) string {
    stack := make([]byte, 0, len(s))
    for i := range s {
	if s[i] == '*' {
	    n := len(stack)
	    if n != 0 {
	        stack = stack[:n-1]
	    }
	} else {
		stack = append(stack, s[i])
	}
    }
    return string(stack)
}
```
