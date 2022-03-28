/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    variant of path-sum 2
    - in this case, we return whether we found a root-to-leaf path that adds up to target
    - once we found a path, we can exit the recursion early.
    
    - we will subtract each node val from targetSum until targetSum == 0
    - this targetSum will act as our local state at each node for each paused recursion call
    - So we wont have to maintain another local state
    
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
    return inorderDfs(root, targetSum, 0)
}

func inorderDfs(root *TreeNode, targetSum, rs int) bool {
    
    // base
    if root == nil {
        return false
    }
    
    // logic
    rs = rs + root.Val
    
    left := inorderDfs(root.Left, targetSum, rs)
    if left { return true }
    
    if rs == targetSum && root.Left == nil && root.Right == nil {
        return true
    }
    
    return inorderDfs(root.Right, targetSum, rs)
    
}