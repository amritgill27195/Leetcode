/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var dfs func(r *TreeNode) int
    sum := 0
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {
            return 0
        }
        
        left := dfs(r.Left)
        right := dfs(r.Right)
        if left+right > sum {
            sum = left+right
        }
        return int(math.Max(float64(left), float64(right)))+1
    }
    dfs(root)
    return sum
}