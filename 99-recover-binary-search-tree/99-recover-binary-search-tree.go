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
        
        if prev != nil && prev.Val >= r.Val {
            // invalid bst
            if first == nil {
                first = prev
            }
            second = r

        }
        prev = r
        
        inorder(r.Right)
    } 
    inorder(root)
    first.Val, second.Val = second.Val, first.Val
}