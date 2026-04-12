# 605. Can Place Flowers

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/can-place-flowers

---
## 🧩 Problem
Given array of integers, return number of flowers that can be planted.
- flowers cannot be planted in neighborhood
- `flowerbed[i]` contains only `0` and `1`.

Constraints:
- `1 <= flowerbed.length <= 2*10^4`
- `0 <= n <= flowerbed.length`
---

## 🔍 Approach


## ⏱ Complexity

- **Time:** `O(n)` – each element is processed once
- **Space:** `O(1)` – only constant space
---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory             |
|----------|----------------|--------------------|
| Rust     | ~0 ms (100%)   | ~2.19  MB (98.96%) |

## 💻 Implementations
Solution iterate through the array and try to plant flowers greedily wherever possible. 
At each position, we check if the current plot and its neighbors are empty; 
if so, we plant a flower and decrease n. 
After planting or encountering an existing flower, we skip the next position since adjacent planting is not allowed. 
If we manage to place all n flowers, we return true; otherwise, false.
### 🦀Rust
```rust
const PLANTED: i32 = 1;

pub fn can_place_flowers(mut flowerbed: Vec<i32>,mut n: i32) -> bool {
    let len = flowerbed.len();
    let mut index = 0;
	if n == 0 { return true; }
    while index < len {
        let has_prev_flower = index > 0 && flowerbed[index - 1] == PLANTED;
        let has_next_flower = index < len -1 && flowerbed[index + 1] == PLANTED;

        if flowerbed[index] == PLANTED {
            index += 2;
            continue;
        }
        if has_prev_flower || has_next_flower {
            index += 1;
            continue;
        }
        flowerbed[index] = PLANTED;
        index += 2;
        n-=1;
        if n <= 0 {
            return true;
        }
    }
    false
}

```