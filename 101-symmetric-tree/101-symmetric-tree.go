/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    We just need a helper that takes 2 trees
    and that helper will check whether
    the 2 trees are mirror image or not.
*/

func isSymmetric(root *TreeNode) bool {
    return dfs(root, root)
}

func dfs(a *TreeNode, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil || a.Val != b.Val {return false}
    
    return dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
}