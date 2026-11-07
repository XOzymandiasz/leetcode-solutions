# 14. Longest Common Prefix
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/longest-common-prefix/

---

## 🧩 Problem

Given a non-empty array of strings, return the longest prefix shared by all of them. If they have no common prefix, return an empty string.

For example, `["flower", "flow", "flight"]` returns `"fl"`, while `["dog", "racecar", "car"]` returns `""`.

Constraints:

- The input contains at least one string.
- Individual strings may be empty.
- Non-empty strings contain lowercase English letters.

---

## 🔍 Approach

Compare strings character by character, using the first string as the reference.

At each index, obtain the reference character with `chars().nth(idx)`. If the first string has no character at that index, the prefix is complete. Otherwise, compare `Some(c)` with the character at the same position in every string. A different character or `None` means that the common prefix ends before this position.

If all strings match, append the character to `answer` with `push` and move to the next index. Character comparisons are case-sensitive.

---

## ⏱ Complexity

Let `n` be the number of strings and `m` the maximum string length.

- **Time:** `O(n · m²)` in the worst case. Each call to `chars().nth(idx)` creates a new iterator and scans from the beginning. Summing this work across positions gives a quadratic factor in `m`.
- **Space:** `O(m)` for the returned prefix, with `O(1)` auxiliary space excluding the result.

---

## ⚡ Performance

> Results are from the supplied LeetCode submission screenshot.
> Runtime and memory measurements depend on the platform's runtime environment.

| Language | Runtime | Memory |
|----------|---------|--------|
| Rust | 0 ms (beats 100.00%) | 2.10 MB (beats 85.90%) |

---

## 💻 Implementations

### 🦀 Rust

[Source](./rust/src/lib.rs) · [Tests](./rust/tests/solution_test.rs)

```rust
pub struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut answer = String::from("");
        let mut idx = 0;
        loop {
            let Some(c) = strs[0].chars().nth(idx) else {
                break;
            };
            let mut not_match = false;
            for str in strs.iter() {
                if str.chars().nth(idx) != Some(c) {
                    not_match = true;
                }
            }
            if not_match {
                break;
            }
            answer.push(c);
            idx += 1;
        }
        answer
    }
}
```
