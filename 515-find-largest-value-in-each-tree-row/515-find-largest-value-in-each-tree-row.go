/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    
    var result []int
    
    var dfs func(level int, a *TreeNode)
    dfs = func(level int, a *TreeNode) {
        if a == nil {
            return
        }
        if len(result) == level {
            result = append(result, a.Val)
        } else {
            if a.Val > result[level] {
                result[level] = a.Val
            }
        }
        dfs(level+1, a.Left)
        dfs(level+1,a.Right)
    }
    
    dfs(0, root)
    return result
    
}