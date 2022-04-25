/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: brute force
    - Find paths that lead to p
    - Find paths that lead to q
    - Find the last common node between the 2 paths
    - The last common node is the lowest common ancestor
    
    time: o(n)
        - worse case our recursion goes left first and p and q are located in the right subtree as leaf nodes
    space: 
        - o(h) to find p + o(h) to find q + o(h) path to p + o(h) path to q
        - o(h)
*/

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     var dfs func(r *TreeNode, target *TreeNode, paths []*TreeNode) (bool,[]*TreeNode)
     dfs = func(r *TreeNode, target *TreeNode, paths []*TreeNode) (bool, []*TreeNode) {
         // base
         if r == nil {
             return false, nil
         }
         
         // logic
         paths = append(paths, r)
         if r.Val == target.Val {
             return true, paths
         }
         foundInLeft, left := dfs(r.Left, target, paths)
         if foundInLeft {
             return true, left
         }
         return dfs(r.Right, target, paths)
     }
     _, pPaths := dfs(root, p, nil)
     _, qPaths := dfs(root, q, nil)
     var out *TreeNode
     pPtr := 0
     qPtr := 0
     for pPtr < len(pPaths) && qPtr < len(qPaths) {
         if pPaths[pPtr] == qPaths[qPtr] {
             out = pPaths[pPtr]
         }
         pPtr++
         qPtr++
     }
     return out
}