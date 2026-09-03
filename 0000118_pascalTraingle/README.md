# 118. Pascal's Triangle
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/pascals-triangle/

---

## 🧩 Problem

Given an integer `numRows`, return the first `numRows` rows of Pascal's triangle.

Each row starts and ends with `1`. Every other value is the sum of the two values directly above it in the previous row.

Example:

```text
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

Constraints:

- `1 <= numRows <= 30`

---

## 🔍 Approach

Start the triangle with its first row: `[1]`.

For every following row:

1. Add `1` as the first value.
2. Calculate each inner value by adding two adjacent values from the previous row.
3. Add `1` as the last value.
4. Append the completed row to the triangle.

---

## ⏱ Complexity

- **Time:** `O(numRows²)` – every value in the triangle is calculated once.
- **Space:** `O(numRows²)` – the returned triangle contains this many values; excluding the result, the algorithm uses `O(1)` auxiliary space.

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory              |
|----------|----------------|---------------------|
| C#       | ~1 ms (63.86%) | ~40.29  MB (68.29%) |

---

## 💻 Implementations

### C#

```csharp
namespace Solution;

public class Solution
{
    public IList<IList<int>> Generate(int numRows)
    {
        var triangle = new List<IList<int>>
        {
            new List<int> { 1 }
        };

        for (var i = 1; i < numRows; i++)
        {
            var row = new List<int>();
            var previousRow = triangle[i - 1];

            row.Add(1);

            for (var j = 0; j < i - 1; j++)
            {
                var value = previousRow[j] + previousRow[j + 1];
                row.Add(value);
            }

            row.Add(1);
            triangle.Add(row);
        }

        return triangle;
    }
}
```
