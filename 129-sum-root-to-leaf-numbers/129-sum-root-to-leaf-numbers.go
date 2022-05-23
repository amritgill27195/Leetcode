/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    sum := 0
    var dfs func(num int, a *TreeNode) 
    dfs = func(num int, a *TreeNode)  {
        // base
        if a == nil {return}
        
        // logic
        num = num * 10 + a.Val
        dfs(num,a.Left)
        dfs(num,a.Right)
        if a.Left == nil && a.Right == nil {
            sum += num
        }
    }
    dfs(0, root)
    return sum
}