# 47. Permutations II

[All Solutions](../)

Difficulty: Medium

https://leetcode.com/problems/permutations-ii/


## Problem

Given an array of integers that may contain duplicates, return all unique permutations in any order.

For example, `[1, 1, 2]` produces `[[1, 1, 2], [1, 2, 1], [2, 1, 1]]`.


## Approach

Sort the input to put equal values next to each other. Backtracking builds one permutation at a time in `row`. The `used` array records which input indices are already in that permutation.

Skip indices that are already used. For equal values, skip the current element if the previous index is still unused:

```rust
if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
    continue;
}
```

Equal elements are therefore selected in index order, so each branch produces a different permutation. For `[1, 1, 2]`, selecting the second `1` requires the first to already be in `row`.

For each choice, mark the index as used and append its value to `row` before recursing. When the call returns, undo that choice: remove the value and mark the index as unused. Once `row` contains all elements, clone it into `answer`.


## Complexity

Let `n` be the number of input elements.

- Time: `O(n × n!)` in the worst case, when all elements are distinct. Sorting costs `O(n log n)`; duplicate pruning reduces the number of branches when values repeat.
- Space: `O(n)` auxiliary space for the recursion stack, `row`, and `used`. Storing the results requires an additional `O(n × U)` space, where `U` is the number of unique permutations.


## Performance

LeetCode submission results (from the screenshot). Runtime and memory measurements depend on the platform's runtime environment.

| Language | Runtime | Memory |
| -------- | ------- | ------ |
| Rust | 0 ms (beats 100.00%) | 2.15 MB (beats 88.46%) |


## Implementations

### Rust

See [source](./rust/src/lib.rs) and [tests](./rust/tests/solution_test.rs).

```rust
pub struct Solution;
impl Solution {
    pub fn permute_unique(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();

        let mut answer = Vec::new();
        let mut row = Vec::with_capacity(nums.len());
        let mut used = vec![false; nums.len()];

        Self::backtrack(&nums, &mut used, &mut row, &mut answer);
        answer
    }

    fn backtrack(nums: &[i32], used: &mut [bool], row: &mut Vec<i32>, answer: &mut Vec<Vec<i32>>) {
        if row.len() == nums.len() {
            answer.push(row.clone());
            return;
        }

        for i in 0..nums.len() {
            if used[i] {
                continue;
            }

            if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
                continue;
            }

            used[i] = true;
            row.push(nums[i]);

            Self::backtrack(nums, used, row, answer);

            row.pop();
            used[i] = false;
        }
    }
}
```

## Tests

Run from the problem directory:

```bash
cargo test --manifest-path rust/Cargo.toml
```

The tests cover three distinct elements, a single element, and duplicate elements (`[1, 1, 2]`). Results are sorted before comparison so permutation order does not affect the tests.
