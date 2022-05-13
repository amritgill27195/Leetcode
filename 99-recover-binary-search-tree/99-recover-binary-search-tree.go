/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    var first *TreeNode
    var second *TreeNode
    var prev *TreeNode
    
    var inorder func(r *TreeNode)
    inorder = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }    
        // logic
        inorder(r.Left)
        if prev != nil {
            if prev.Val >= r.Val {
                if second == nil {
                    second = prev
                }
                first = r
            }
        }
        prev = r
        inorder(r.Right)
    }
    inorder(root)
    first.Val, second.Val = second.Val, first.Val
    
}