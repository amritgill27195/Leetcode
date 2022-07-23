/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode{
        // base
        
        // logic
        if r.Val == p.Val || r.Val == q.Val {
            return r
        }
        if p.Val > r.Val && q.Val > r.Val {
            return dfs(r.Right)
        }
        if p.Val < r.Val && q.Val < r.Val {
            return dfs(r.Left)
        }
        return r
    }
    return dfs(root)
}