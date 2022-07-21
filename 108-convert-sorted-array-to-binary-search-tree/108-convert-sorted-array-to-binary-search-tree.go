/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    var dfs func(n []int) *TreeNode
    dfs = func(n []int) *TreeNode {
        // base
        if n == nil || len(n) == 0 {return nil}
        
        // logic
        rootIdx := len(n) / 2
        root := &TreeNode{Val: n[rootIdx]}
        root.Left = dfs(n[:rootIdx])
        root.Right = dfs(n[rootIdx+1:])
        return root
    }
    
    return dfs(nums)
}