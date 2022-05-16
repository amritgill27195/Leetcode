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
    var dfs func(r *TreeNode, path []int, target int)
    dfs = func(r *TreeNode, path []int, target int) {
        // base
        if r == nil {return}
       
        // logic
        path = append(path, r.Val)
        target = target-r.Val

        dfs(r.Left, path, target)
        
        if target == 0 && r.Left == nil && r.Right == nil {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return
        }
        
        dfs(r.Right, path, target)
        path = path[:len(path)-1]
        
    }
    dfs(root, nil, targetSum)
    return result
}