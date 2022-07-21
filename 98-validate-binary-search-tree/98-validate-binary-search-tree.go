/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var dfs func(min, max int, r *TreeNode) bool
    dfs = func(min, max int, r *TreeNode) bool {
        // base
        if r == nil {return true}
        // logic
        if r.Val <= min || r.Val >= max {
            return false
        }
        return dfs(min, r.Val, r.Left) && dfs(r.Val, max, r.Right)
    }
    return dfs(math.MinInt64, math.MaxInt64, root)
}