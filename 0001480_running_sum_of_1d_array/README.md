# 1480. Running Sum of 1d Array

[All Solutions](../)

Difficulty: Easy

https://leetcode.com/problems/running-sum-of-1d-array/

## Problem

Return the running sum of an integer array. Each element in the result is the sum of the input elements from index `0` through that position.

For example, `[1, 2, 3, 4]` becomes `[1, 3, 6, 10]`.

## Approach

Update the input vector in place, starting at index `1`. Add the previous element to the current one:

```rust
nums[i] += nums[i - 1]
```

The previous element already holds the sum up to that point. Adding it to the current value gives the next running sum. The first element stays the same. Return the updated vector after the loop.

## Complexity

For an input of `n` elements:

- Time: `O(n)`, with one pass through the vector.
- Space: `O(1)` auxiliary space. The function reuses the input vector for the result.

## Performance

LeetCode submission results. Measurements depend on the platform's runtime environment.

| Language | Runtime | Memory |
| -------- | ------- | ------ |
| Rust | 0 ms (beats 100.00%) | 2.13 MB (beats 60.71%) |

## Implementation

### Rust

See [source](./rust/src/lib.rs) and [tests](./rust/tests/solution_test.rs).

```rust
pub struct Solution;
impl Solution {
    pub fn running_sum(mut nums: Vec<i32>) -> Vec<i32> {
        for i in 1..nums.len() {
            nums[i] += nums[i - 1]
        }
        nums
    }
}
```

## Tests

The tests check repeated ones (`[1, 1, 1, 1, 1]`) and increasing values (`[1, 2, 3, 4]`).

Run from the problem directory:

```bash
cargo test --manifest-path rust/Cargo.toml
```
