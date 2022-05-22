/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    max := math.MinInt64
    var dfs func(a *TreeNode) int
    dfs = func(a *TreeNode) int {
        // base
        if a == nil {
            return 0
        }
        
        // logic
        leftSum := int(math.Max(float64(dfs(a.Left)), 0.0))
        rightSum := int(math.Max(float64(dfs(a.Right)), 0.0))
        
        sumWithCurrentNode := a.Val + leftSum + rightSum
        max = int(math.Max(float64(max), float64(sumWithCurrentNode)))
        
        return a.Val + int(math.Max(float64(leftSum),float64(rightSum)))
    }
    dfs(root)
    return max
}