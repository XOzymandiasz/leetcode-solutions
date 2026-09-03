# 119. Pascal's Triangle II
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/pascals-triangle-ii/

---

## 🧩 Problem

Given an integer `rowIndex`, return the `rowIndex`-th row of Pascal's triangle, using zero-based indexing.

Each row starts and ends with `1`. Every other value is the sum of the two values directly above it in the previous row.

Example:

```text
Input: rowIndex = 3
Output: [1,3,3,1]
```

Constraints:

- `0 <= rowIndex <= 33`

---

## 🔍 Approach

The values in row `n` are binomial coefficients:

```text
C(n, 0), C(n, 1), ..., C(n, n)
```

Start with `C(n, 0) = 1` and calculate every following value from the previous one:

```text
C(n, k) = C(n, k - 1) * (n - k + 1) / k
```

This produces the requested row directly, without generating any of the preceding rows. A `long` is used for the intermediate multiplication so it does not overflow before the result is converted to `int`.

---

## ⏱ Complexity

- **Time:** `O(rowIndex)` – each value in the requested row is calculated once.
- **Space:** `O(rowIndex)` – the returned list contains `rowIndex + 1` values; excluding the result, the algorithm uses `O(1)` auxiliary space.

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory              |
|----------|----------------|---------------------|
| Go       | ~0 ms (100%)   | ~39.77  MB (61.39%) |
---

## 💻 Implementations

### C#

```csharp
namespace Solution;

public class Solution
{
    public IList<int> GetRow(int rowIndex)
    {
        var row = new List<int>();

        var maxRowIndex = rowIndex + 1;
        long currentNumber = 1;
        row.Add((int)currentNumber);

        for (var i = 1; i < maxRowIndex; i++)
        {
            currentNumber *= maxRowIndex - i;
            currentNumber /= i;
            row.Add((int)currentNumber);
        }

        return row;
    }
}
```
