# 2. Add Two Numbers

🟠 Difficulty: Medium

🔗 https://leetcode.com/problems/add-two-numbers/

---
## 🧩 Problem
Given two linked lists `l1` and `l2`, return their sum as a `linked list`. 
- Digits are stored in reverse order
- Digits are non-negative
- If both numbers are empty, return 0

Constraints:
- The number of nodes is in the range [0, 100]
- 0 <= Node.val <= 9

ListNode:
```
ListNode:
    value: integer
    next: reference to ListNode
```
## 🔍 Approach
We iterate through both linked lists simultaneously, adding corresponding digits.

- At each step, we sum the current node values and the carry
- A new node is created to store `sum % 10`
- The carry (`sum / 10`) is propagated to the next iteration
- We continue until both lists are fully processed and no carry remains

The result list is constructed incrementally during the iteration.

---

## ⏱ Complexity

- **Time:** `O(max(n, m))`
- **Space:** `O(max(n, m))`

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language    | Runtime        | Memory            |
|-------------|----------------|-------------------|
| Go          | ~0 ms (100%)   | ~6.1  MB (87.38%) |
---
## 💻 Implementations

### 🐹 Go
```go
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    current := &ListNode{}
    result := current
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        current.Val = sum % 10
        carry = sum / 10
        if l1 != nil || l2 != nil || carry != 0 {
            current.Next = &ListNode{}
            current = current.Next
        }
    }
    return result	
}
```