/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    result := 0
    var dfs func(max int, a *TreeNode)
    dfs = func(max int, a *TreeNode) {
        // base
        if a == nil {return}
        
        // logic
        if a.Val >= max {
            result++
        }
        dfs(int(math.Max(float64(max), float64(a.Val))), a.Left)
        dfs(int(math.Max(float64(max), float64(a.Val))), a.Right)
    }
    dfs(root.Val, root)
    return result
}