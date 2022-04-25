/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count := 0
    var inorder func(r *TreeNode, k int) int
    inorder = func(r *TreeNode, k int) int {
        // base
        if r == nil {
            return -1
        }
        
        
        // logic
        left := inorder(r.Left, k)
        if left != -1 {
            return left
        }
        count++
        if count == k {
            return r.Val
        }
        right := inorder(r.Right, k)
        if right != -1 {
            return right
        }
        return -1
    }
    return inorder(root, k)
}