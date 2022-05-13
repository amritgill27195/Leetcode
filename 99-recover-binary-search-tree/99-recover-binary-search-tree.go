/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    
    list := []int{}
    var inorder func(r *TreeNode)
    inorder = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        inorder(r.Left)
        list = append(list, r.Val)
        inorder(r.Right)
    }
    
    inorder(root)
    sort.Ints(list)
    
    lPtr := 0
    var inorderW func(r *TreeNode)
    inorderW = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        inorderW(r.Left)
        r.Val = list[lPtr]
        lPtr++
        inorderW(r.Right)
    }
    
    inorderW(root)
    
}