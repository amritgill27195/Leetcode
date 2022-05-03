/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// simple inorder until you hit p, once you hit p, return the next node after p
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {
            return nil
        }
        
        left := dfs(r.Left)
        if left != nil {
            return left
        }
        if r.Val > p.Val {
            return r
        }
        return dfs(r.Right)
    } 
    return dfs(root)
}