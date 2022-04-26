/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    max := 0
    var dfs func(r *TreeNode) int 
    dfs = func(r *TreeNode) int {
        // bottom up recursion 
        // child returns its height back to parent
        // parent gets height from both childs
        // adds them up
        // and saves it if sum > max
        // and then returns the max back to its parent
        
        // base
        if r == nil {
            return 0
        }
        
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        cmax := left+right
        if cmax > max {
            max = cmax
        }
        
        return int(math.Max(float64(left), float64(right)))+1
        
    }
    dfs(root)
    return max
}