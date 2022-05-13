/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var inorder func(r *TreeNode) int
    inorder = func(r *TreeNode) int {
        // base
        if r == nil {
            return -1
        }
        
        // logic
        left := inorder(r.Left)
        if left != -1 {
            return left
        }
        
        k--
        if k == 0 {
            return r.Val
        }
        
        return inorder(r.Right)
    }
    return inorder(root)
}