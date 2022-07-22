/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(a, b *TreeNode) bool
    dfs = func(a, b *TreeNode) bool {
        // base
        
        // logic
        if a == nil && b == nil {return true}
        if a == nil || b == nil {return false}
        return a.Val == b.Val && dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
    }
    return dfs(p,q)
}