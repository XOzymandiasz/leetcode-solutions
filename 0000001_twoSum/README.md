# 1. Two Sum

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/two-sum/

---

## 🧩 Problem
Given an array of integers `nums` and an integer `target`, returns indices of two numbers whose sum equals the `target`.

Constraints:
- Each input has exactly one solution.
- The answer can be returned in any order.

---

## 🔍 Approach
The solution uses a **map** to track previously visited numbers and their indices.

### Key ideas

- For each element `n`, compute its complement: `target - n`
- Check if the complement has already been seen
- If yes — return both indices immediately
- Otherwise — store the current number in the map

### Why it works

Instead of checking all pairs (which would take `O(n²)`), the hash map allows us to:
- store values we've seen
- check complements in constant time `O(1)`

This reduces the overall complexity to linear time.
---

## ⏱ Complexity

- **Time:** `O(n)` - single pass through the array
- **Space:** `O(n)` - additional memory for the hash map

---
## ⚡ Performance

| Language    | Runtime        | Memory            |
|-------------|----------------|-------------------|
| Rust        | ~0 ms (100%)   | ~2.3  MB (83.34%) |
| Go          | ~0 ms (100%)   | ~5.8  MB (71.28%) |
| C++         | ~0 ms (100%)   | ~14.7 MB (52.32%) |
| Python      | ~0 ms (100%)   | ~20.4 MB (64.31%) | 
| PHP         | ~0 ms (100%)   | ~20.6 MB (79.88%) | 
| JavaScript  | ~0 ms (100%)   | ~56.4 MB (56.45%) | 
| TypesScript | ~0 ms (100%)   | ~56.8 MB (59.13%) | 
| Java        | ~2 ms (99.23%) | ~47.1 MB (68.47%) | 
---
## 💻 Implementations

## 🦀Rust
``` rust
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    use std::collections::HashMap;
    let mut visited: HashMap<i32, i32> = HashMap::with_capacity(nums.len());

    for (i, &n) in nums.iter().enumerate() {
        if let Some(&idx) = visited.get(&(target - n)) {
            return vec![idx, i as i32];
        }
        visited.insert(n, i as i32);
    }
    vec!{}
}
```
## 🐹 Go
```go
func TwoSum(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))

	for i, n := range nums {
		if idx, exist := visited[target-n]; exist {
			return []int{idx, i}
		}

		visited[n] = i
	}

	return []int{}
}
```
## 🔧C
```c
typedef struct  {
    int key;
    int value;
    UT_hash_handle hh;
} HashItem;

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    HashItem* map = NULL;

    for (int i = 0; i < numsSize; ++i) {
        int complement = target - nums[i];
        
        HashItem* item;
        HASH_FIND_INT(map, &complement, item);
        
        if (item != NULL) {
            int* result = malloc(2 * sizeof(int));
            result[0] = item->value;
            result[1] = i;
            *returnSize = 2;
            return result;
        }
        
        item = malloc(sizeof(HashItem));
        item->key = nums[i];
        item->value = i;
        HASH_ADD_INT(map, key, item);
    }

    
    returnSize = NULL;
    return NULL;
}
```

## ⚙️C++
```cpp
std::vector<int> twoSum(std::vector<int>& nums, int target) {
    std::unordered_map<int, int> m;

    for (size_t i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (m.find(complement) != m.end()) {
            return {m[complement], static_cast<int>(i)};
        }

        m[nums[i]] = i;
    }

    return {};
}
```

## 🐍 Python
```python
def twoSum(self, nums: List[int], target: int) -> List[int]:
    visited = {}
    
    for i, num in enumerate(nums):
        if target - num in visited:
            return [visited[target - num], i]
        visited[num] = i

    return []
```
## 🟪🐘 PHP
```php
/**
 * @param Integer[] $nums
 * @return Integer[]
 */
function twoSum(array $nums, int $target): array {
    $visited = [];
    foreach ($nums as $i => $num) {
        if (isset($visited[$target - $num])) {
            return [$visited[$target - $num], $i];
        }
        $visited[$num] = $i;     
    }
    return [];
}
```
## 🟨 Javascript
```javascript
var twoSum = function(nums, target) {
    const visited = new Map();

    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];

        if (visited.has(complement)) {
            return [visited.get(complement), i];
        }

        visited.set(nums[i], i);
    }

    return [];
};
```
## 🔷 Typescript
```typescript
function twoSum(nums: number[], target: number): number[] {
    const visited: Record<number, number> = {};
    
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (complement in visited) {
            return [visited[complement], i];
        }
        visited[nums[i]] = i;
    }
    
    return []
}
```

## ☕ Java
```java
public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
        Integer complement = target - nums[i];
        Integer previous = map.get(complement);

        if (previous != null) {
            return new int[]{previous, i};
        }

        map.put(nums[i], i);
    }
    return new int[0];
}
```