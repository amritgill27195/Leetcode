/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// just traverse with bst constraints 
func rangeSumBST(root *TreeNode, low int, high int) int {
    b := &bst{sum: 0}
    b.inorderDfs(root, low, high)
    return b.sum
}

type bst struct {
    sum int
}
func (b *bst)inorderDfs(root *TreeNode, low, high int) {
    // base 
    if root == nil || (root.Val > high && root.Val < low) {
        return
    }
    
    
    b.inorderDfs(root.Left, low, high)
    
    
    if root.Val >= low && root.Val <= high {
        b.sum += root.Val
    }
    
    
    b.inorderDfs(root.Right, low, high)
}