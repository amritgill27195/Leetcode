/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool) {
        // base
        if r == nil {return 0, true}
        
        // logic
        left, leftBalanced := dfs(r.Left)
        right, rightBalanced := dfs(r.Right)
        
        if !leftBalanced || !rightBalanced {
            return -1, false
        }
        if abs(left-right) > 1 {
            return -1, false
        }
        return max(left, right)+1, true 
        
    }
    _, res := dfs(root)
    return res
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}