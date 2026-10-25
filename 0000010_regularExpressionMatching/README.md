# 10. Regular Expression Matching
[All Solutions](../)

🔴 Difficulty: Hard

🔗 https://leetcode.com/problems/regular-expression-matching/

---

## 🧩 Problem

Determine whether the entire string `s` matches the pattern `p`.

- `.` matches any single character.
- `*` matches zero or more occurrences of the preceding character.
- Other characters must match literally.

For example, `"aa"` does not match `"a"`, but matches `"a*"`. The pattern `".*"` matches `"ab"`.

Each `*` in the pattern has a valid preceding character to repeat.

---

## 🔍 Approach

Use recursion with memoization. State `(i, j)` describes whether the text starting at index `i` matches the pattern starting at index `j`.

Check whether the current characters match literally or through `.`. When the next pattern character is `*`, try skipping the character and its star (zero repetitions). Alternatively, if the current characters match, consume one text character and keep the same pattern position to allow more repetitions.

Without a following `*`, consume one character from both strings if they match. When the pattern ends, the text must also have ended.

Cache each result in `memo` so every state is evaluated at most once. Character vectors allow indexed access to both strings.

---

## ⏱ Complexity

Let `n` be the text length and `m` the pattern length in characters.

- **Time:** `O((n + 1)(m + 1))` to initialize the memo table and evaluate its states, each in constant work.
- **Space:** `O((n + 1)(m + 1))` for memoization, plus `O(n + m)` for character vectors and the recursion stack.

For nonempty inputs, both bounds simplify to `O(nm)`.

---

## ⚡ Performance

> Results are based on the LeetCode submission shown in the provided screenshot.
> Percentiles depend on the platform's runtime environment and are not rigorous benchmarks.

| Language | Runtime | Memory |
|----------|---------|--------|
| Rust | 0 ms (100.00%) | 2.05 MB (92.91%) |

---

## 💻 Implementations

### 🦀 Rust

[Source](./rust/src/lib.rs) · [Tests](./rust/tests/solution_test.rs)

```rust
pub struct Solution;
impl Solution {
    pub fn is_match(s: String, p: String) -> bool {
        let text: Vec<char> = s.chars().collect();
        let pattern: Vec<char> = p.chars().collect();
        let mut memo = vec![vec![None; pattern.len() + 1]; text.len() + 1];

        Self::matches_from(&text, &pattern, 0, 0, &mut memo)
    }

    fn matches_from(
        text: &[char],
        pattern: &[char],
        i: usize,
        j: usize,
        memo: &mut [Vec<Option<bool>>],
    ) -> bool {
        if let Some(result) = memo[i][j] {
            return result;
        }

        let result = if j == pattern.len() {
            i == text.len()
        } else {
            let first_matches = i < text.len() && (pattern[j] == '.' || pattern[j] == text[i]);

            if pattern.get(j + 1) == Some(&'*') {
                Self::matches_from(text, pattern, i, j + 2, memo)
                    || (first_matches && Self::matches_from(text, pattern, i + 1, j, memo))
            } else {
                first_matches && Self::matches_from(text, pattern, i + 1, j + 1, memo)
            }
        };

        memo[i][j] = Some(result);
        result
    }
}
```

