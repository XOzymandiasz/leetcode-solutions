# 700. Search In a Binary Search Tree
[All Solutions](../)  

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/search-in-a-binary-search-tree/

---
## 🧩 Problem
Given the root of a binary search tree and an integer val, return the subtree rooted at the node with value val. 
- If such a node does not exist, return nil.

ListNode:
```
ListNode:
    value: integer
    left  reference to ListNode
    right reference to ListNode
```

Constraints:
- `1 <= Node.val <= 10^7`
- `1 <= val <= 10^7`
- `1 <= h (tree height) <= 5000`

---

## 🔍 Approach
We take advantage of the BST property. 
Starting from the root, if the current node’s value equals val, we return it. 
If val is greater, we move to the right subtree; otherwise, we move to the left subtree. 
We continue this process recursively (or iteratively) until we find the node or reach nil.
---

## ⏱ Complexity

- **Time:** `O(h)` - height of the tree
- **Space:** `O(h)` - due to recursion

---
## ⚡ Performance
> Results are based on LeetCode submissions for the same algorithmic approach.  
> Values are approximate and depend on the platform's runtime environment, so they should not be treated as rigorous benchmarks.

| Language | Runtime        | Memory            |
|----------|----------------|-------------------|
| Go       | ~0 ms (100%)   | ~8.5  MB (99.51%) |

## 💻 Implementations

### 🐹 Go
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    if root.Val < val {
        return SearchBST(root.Right, val)
    }

    return SearchBST(root.Left, val)
}
```