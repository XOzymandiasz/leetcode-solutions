# 66. Plus One
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/plus-one/

---

## 🧩 Problem

Given a non-negative integer represented as an array of digits, add one and return the resulting digits. The digits are ordered from most significant to least significant.

For example, `[1, 2, 3]` becomes `[1, 2, 4]`, and `[9, 9]` becomes `[1, 0, 0]`.

Constraints:

- The array contains at least one digit.
- Each digit is between `0` and `9`.
- The number has no leading zeros, except for zero itself.

---

## 🔍 Approach

Walk through the digits from right to left. If a digit is less than `9`, increment it and return the vector immediately.

A `9` becomes `0`, and the carry continues to the next digit. If every digit was `9`, insert `1` at the beginning. For example, `[9, 9]` first becomes `[0, 0]`, then `[1, 0, 0]`.

The function takes ownership of the vector and updates its elements without cloning it.

---

## ⏱ Complexity

- **Time:** `O(n)` in the worst case, when every digit is `9`. If the last digit is less than `9`, the function returns in `O(1)` time.
- **Space:** `O(1)` auxiliary space. The result reuses the input vector, but inserting a leading `1` may require an `O(n)` allocation if the vector has no spare capacity.

---

## ⚡ Performance

> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime | Memory |
|----------|---------|--------|
| Rust | ~0 ms (100.00%) | ~2.17 MB (59.55%) |

---

## 💻 Implementations

### 🦀 Rust

```rust
pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
    let mut digits = digits;

    for digit in digits.iter_mut().rev() {
        if *digit < 9 {
            *digit += 1;
            return digits;
        }
        *digit = 0
    }

    digits.insert(0, 1);
    digits
}
```
