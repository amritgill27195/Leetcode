/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    
    
    var dfs func(a *TreeNode) *TreeNode
    dfs = func(a *TreeNode) *TreeNode{
        // base
        if a == nil {
            return nil
        }
        
        // logic
        left := dfs(a.Left)
        right := dfs(a.Right)
        
        if a.Left != nil {
            left.Right = a.Right
            a.Right = a.Left
            a.Left = nil
        }
        if right != nil {return right}
        if left != nil {return left}
        return a
    }
    dfs(root)
}