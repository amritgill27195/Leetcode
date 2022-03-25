/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {    
    return inorder(root, 0)
}


func inorder(root *TreeNode, runningNum int) int  {
    // base
    if root == nil {
        return 0
    }
    
    runningNum = runningNum * 10 + root.Val
    // logic
    left := inorder(root.Left, runningNum)
    if root.Left == nil && root.Right == nil {
        return runningNum
    }
    
    right := inorder(root.Right, runningNum)
    if root.Left == nil && root.Right == nil {
        return runningNum
    }
    return left + right
    
}