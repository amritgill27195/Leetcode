/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    sum := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        if r.Val > low {
            dfs(r.Left)
        }
        if r.Val < high {
            dfs(r.Right)
        }
        if r.Val >= low && r.Val <= high {
            sum += r.Val
        }
    }
    dfs(root)
    return sum
}