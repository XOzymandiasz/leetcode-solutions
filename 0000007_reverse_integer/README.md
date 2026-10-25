# 7. Reverse Integer
[All Solutions](../)  

🟡 Difficulty: Medium

🔗 https://leetcode.com/problems/reverse-integer/

---

## 🧩 Problem

Reverse the decimal digits of a signed 32-bit integer. Return `0` if the reversed value falls outside the signed 32-bit range.

For example, `123` becomes `321`, `-123` becomes `-321`, and `120` becomes `21`. Reversing `1534236469` would exceed the range, so the result is `0`.

Constraints:

- The input is between `-2^31` and `2^31 - 1`.
- The solution must not store values in a 64-bit integer.

---

## 🔍 Approach

Extract the last digit with `x_copy % 10`, then append it to the answer using `answer * 10 + digit`. Divide `x_copy` by `10` to remove the processed digit and repeat until no digits remain.

Rust's integer division truncates toward zero, and the remainder keeps the dividend's sign. The same loop therefore handles negative inputs without taking an absolute value.

Use `checked_mul` and `checked_add` to detect overflow before storing the next answer. If either operation returns `None`, return `0`. An input of zero skips the loop and returns zero.

---

## ⏱ Complexity

- **Time:** `O(d)`, where `d` is the number of decimal digits. Each iteration processes one digit, with at most 10 iterations for an `i32`.
- **Space:** `O(1)`. The function uses a fixed number of integer variables.

---

## ⚡ Performance

> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime | Memory |
|----------|---------|--------|
| Rust | ~0 ms (100.00%) | ~2.16 MB (45.85%) |

---

## 💻 Implementations

### 🦀 Rust

```rust
pub fn reverse(x: i32) -> i32 {
    let mut x_copy = x;
    let mut answer: i32 = 0;

    while x_copy != 0 {
        let digit = x_copy % 10;

        answer = match answer
            .checked_mul(10)
            .and_then(|value| value.checked_add(digit))
        {
            Some(value) => value,
            None => return 0,
        };

        x_copy /= 10;
    }

    answer
}
```
