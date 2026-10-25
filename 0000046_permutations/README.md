# 46. Permutations

[All Solutions](../)

🟡 Difficulty: Medium

🔗 https://leetcode.com/problems/permutations/

---

## 🧩 Problem

Given an array of distinct integers, return all possible permutations. The permutations can be returned in any order.

For example, `[1, 2, 3]` produces `[[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]`.

Constraints:

- `1 <= nums.length <= 6`.
- `-10 <= nums[i] <= 10`.
- All integers in `nums` are unique.

---

## 🔍 Approach

Build each permutation recursively. At each step, choose one of the remaining numbers, remove it from a cloned vector, and append it to a copy of the current row. Continue recursively with the remaining numbers.

When no numbers remain, append the completed row to `answer`. Passing `answer` as `&mut Vec<Vec<i32>>` lets every recursive call write to the same result vector.

Each branch uses its own copies of the remaining numbers and the current row, so no explicit undo step is needed.

---

## ⏱ Complexity

Let `n` be the number of input elements.

- **Time:** `O(n × n!)`. There are `n!` permutations, each containing `n` elements. Cloning vectors and removing elements also contribute to the work at each recursive step.
- **Space:** `O(n²)` auxiliary space, because recursive calls retain their own vectors along a path of depth `n`. Storing all results requires an additional `O(n × n!)` space.

---

## ⚡ Performance

> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime         | Memory            |
| -------- | --------------- | ----------------- |
| Rust     | ~0 ms (100.00%) | ~2.14 MB (84.21%) |

---

## 💻 Implementations

### 🦀 Rust

See [source](./rust/src/lib.rs) and [tests](./rust/tests/solution_test.rs).

```rust
pub struct Solution;
impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut answer: Vec<Vec<i32>> = Vec::new();
        Self::get_permutations(&mut answer, nums, Vec::new());
        answer
    }

    fn get_permutations(answer: &mut Vec<Vec<i32>>, nums: Vec<i32>, row: Vec<i32>) {
        if nums.is_empty() {
            answer.push(row.clone());
            return;
        }
        for i in 0..nums.len() {
            let mut remaining = nums.clone();
            let value = remaining.remove(i);
            let mut next_row: Vec<i32> = row.clone();
            next_row.push(value);
            Self::get_permutations(answer, remaining, next_row);
        }
    }
}
```

## 🧪 Tests

Run from the problem directory:

```bash
cargo test --manifest-path rust/Cargo.toml
```

The tests cover all six permutations of three elements and the single-element case. Results are sorted before comparison so permutation order does not affect the test.
