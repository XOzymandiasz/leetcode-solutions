# 65. Valid Number
[All Solutions](../)  

🔴 Difficulty: Hard

🔗 https://leetcode.com/problems/valid-number/

---

## 🧩 Problem

Return whether a string represents a valid integer or decimal number, optionally followed by an exponent.

The number may begin with `+` or `-`. A decimal point requires at least one digit on either side: `4.`, `.9`, and `3.14` are valid. An exponent starts with `e` or `E` and must contain an integer, optionally signed.

For example, `-0.1`, `2e10`, and `3e+7` are valid. Strings such as `1e`, `--6`, and `99e2.5` are invalid.

Constraints:

- Digits are ASCII characters from `0` to `9`.
- The number must contain at least one digit before any exponent.
- An exponent must contain at least one digit and cannot contain a decimal point.
- Signs are allowed only at the start of the number or immediately after `e` or `E`.

---

## 🔍 Approach

Scan the string from left to right. Constant slices list the punctuation allowed at each stage, while digits are checked separately.

After each digit, update the allowed punctuation using `has_dot` and `has_exponent`. Once a decimal point appears, another point is forbidden. After an exponent begins, only digits and an optional initial sign are allowed.

The `has_digit` flag records whether the current part contains a digit. An exponent requires this flag to be set, then resets it so the exponent must supply its own digits. Returning `has_digit` rejects empty strings, lone signs, and incomplete exponents such as `1e+`.

Any character that is neither a digit nor allowed punctuation causes an immediate `false` result.

---

## ⏱ Complexity

- **Time:** `O(n)`, where `n` is the string length. Each character is checked once, and the constant slices have fixed sizes.
- **Space:** `O(1)` auxiliary space. The function uses a slice reference and three boolean flags.

---

## ⚡ Performance

> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime | Memory |
|----------|---------|--------|
| Rust | ~0 ms (100.00%) | ~2.05 MB (97.67%) |

---

## 💻 Implementations

### 🦀 Rust

```rust
const FIRST_VALID_SIGNS: &[char] = &['-', '+', '.'];
const AFTER_SIGN_VALID_SIGNS: &[char] = &['.'];
const AFTER_NUMBER_VALID_SIGNS: &[char] = &['.', 'e', 'E'];
const AFTER_DOT_VALID_SIGNS: &[char] = &['e', 'E'];
const AFTER_EXPONENT_VALID_SIGNS: &[char] = &['-', '+'];
const NUMBERS: &[char] = &['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];

pub fn is_number(param: String) -> bool {
    let mut valid = FIRST_VALID_SIGNS;
    let mut has_digit = false;
    let mut has_dot = false;
    let mut has_exponent = false;

    for sign in param.chars() {
        if NUMBERS.contains(&sign) {
            has_digit = true;

            valid = if has_exponent {
                &[]
            } else if has_dot {
                AFTER_DOT_VALID_SIGNS
            } else {
                AFTER_NUMBER_VALID_SIGNS
            };

            continue;
        }

        if !valid.contains(&sign) {
            return false;
        }

        match sign {
            '+' | '-' => {
                valid = if has_exponent {
                    &[]
                } else {
                    AFTER_SIGN_VALID_SIGNS
                };
            }
            '.' => {
                has_dot = true;
                valid = AFTER_DOT_VALID_SIGNS;
            }
            'e' | 'E' => {
                if !has_digit {
                    return false;
                }

                has_exponent = true;
                has_digit = false;
                valid = AFTER_EXPONENT_VALID_SIGNS;
            }
            _ => return false,
        }
    }

    has_digit
}
```
