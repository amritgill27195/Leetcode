/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    
    // base
    if root == nil {
        return
    }
    
    //logic
    flatten(root.Left)
    flatten(root.Right)
    if root.Left != nil {
        tmp := root.Right
        root.Right = root.Left
        root.Left = nil
        for root.Right != nil {
            root = root.Right
        }
        root.Right = tmp
    }
    
}