# 69. Sqrt(x)
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/sqrtx/

---

## 🧩 Problem

Given a nonnegative integer `x`, return its square root rounded down to the nearest integer, without using built-in exponent or square root functions.

For example, `4` returns `2`, and `8` also returns `2` because its square root lies between `2` and `3`.

Constraints:

- `0 <= x <= 2^31 - 1`.

---

## 🔍 Approach

Use binary search over the inclusive integer interval `[0, x]`. The variable `y` stores the largest checked value whose square does not exceed `x`.

Calculate `mid = x_1 + (x_2 - x_1) / 2`. If `mid * mid <= x`, save `mid` in `y` and search larger candidates by setting `x_1 = mid + 1`. Otherwise, search smaller candidates by setting `x_2 = mid - 1`.

Stop when the interval is empty. The saved candidate is the square root rounded down. This also handles `0` and `1` without special cases.

Convert the input to `i64` before searching so that squaring a candidate cannot overflow for any valid `i32` input.

---

## ⏱ Complexity

- **Time:** `O(log(x + 1))`, with constant work for `x = 0`. Each iteration halves the search interval.
- **Space:** `O(1)`. Only a fixed number of integer variables is needed.

---

## ⚡ Performance

LeetCode runtime and memory results have not been recorded for this solution.

---

## 💻 Implementations

### 🦀 Rust

[Source](./rust/src/lib.rs) · [Tests](./rust/tests/solution_test.rs)

```rust
pub struct Solution;
impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        let x = i64::from(x);
        let mut x_1 = 0;
        let mut x_2 = x;
        let mut y = 0;

        while x_1 <= x_2 {
            let mid = x_1 + (x_2 - x_1) / 2;

            if mid * mid <= x {
                y = mid;
                x_1 = mid + 1;
            } else {
                x_2 = mid - 1;
            }
        }

        y as i32
    }
}
```

