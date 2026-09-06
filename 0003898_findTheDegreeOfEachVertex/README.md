# 3898. Find the Degree of Each Vertex
[All Solutions](../)

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/find-the-degree-of-each-vertex/

---

## 🧩 Problem

You are given an `n x n` integer array `matrix` representing the adjacency matrix of an undirected graph with `n` vertices labeled from `0` to `n - 1`.

- `matrix[i][j] == 1` means that vertices `i` and `j` are connected by an edge.
- `matrix[i][j] == 0` means that there is no edge between vertices `i` and `j`.

Return an integer array `ans` of length `n`, where `ans[i]` is the degree of vertex `i`.

Examples:

```text
Input: matrix = [[0,1,1],[1,0,1],[1,1,0]]
Output: [2,2,2]

Input: matrix = [[0,1,0],[1,0,0],[0,0,0]]
Output: [1,1,0]

Input: matrix = [[0]]
Output: [0]
```

Constraints:

- `1 <= n == matrix.length == matrix[i].length <= 100`
- `matrix[i][i] == 0`
- `matrix[i][j]` is either `0` or `1`
- `matrix[i][j] == matrix[j][i]`

---

## 🔍 Approach

In an adjacency matrix, every `1` in row `i` represents an edge incident to vertex `i`. Because the graph has no self-loops, the sum of row `i` is therefore exactly the degree of vertex `i`.

Map every row to its sum and collect the sums in the result array.

---

## ⏱ Complexity

- **Time:** `O(n²)` – all `n²` entries of the adjacency matrix are visited once.
- **Space:** `O(n)` – the result array contains one degree for each vertex; auxiliary space is `O(1)`.

---

## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory              |
|----------|----------------|---------------------|
| C#       | ~13 ms (6.58%) | ~61.69  MB (22.37%) |

---

## 💻 Implementations

### C#

```csharp
namespace Solution;

public class Solution
{
    public int[] FindDegrees(int[][] matrix) 
    {
        return matrix
            .Select(row => row.Sum())
            .ToArray();
    }
}
```
