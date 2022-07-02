/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    
    result := [][]int{}
    
    var dfs func(sum int, r *TreeNode, path []int)
    dfs = func(sum int, r *TreeNode, path []int) {
        // base
        if r == nil {
            return
        }
        path = append(path, r.Val)
        sum += r.Val        
        if sum == targetSum && r.Left == nil && r.Right == nil {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return
        }
        
        
        // logic
        dfs(sum, r.Left, path)
        dfs(sum, r.Right, path)
        path = path[:len(path)-1]
    }
    
    dfs(0, root, nil)
    return result
    
}