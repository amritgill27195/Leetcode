/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}
        
        left := dfs(r.Left)
        right := dfs(r.Right)
        
        if left && right {
            
            if r.Left == nil && r.Right == nil {
                count++
                return true
            } else if r.Left != nil && r.Right != nil {
                if r.Left.Val == r.Val && r.Val == r.Right.Val {
                    count++
                    return true
                }
            } else if r.Left == nil {
                if r.Val == r.Right.Val {
                    count++
                    return true
                }
            } else if r.Right == nil {
                if r.Val == r.Left.Val {
                    count++
                    return true
                }
            }
            
        }
        return false
        
    }
    dfs(root)
    return count
}