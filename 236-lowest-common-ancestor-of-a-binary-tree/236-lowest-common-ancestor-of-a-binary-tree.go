/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var pPaths []*TreeNode
    var qPaths []*TreeNode

    var dfs func(r *TreeNode, paths []*TreeNode)
    dfs = func(r *TreeNode, paths []*TreeNode)  {
        // base
        if r == nil {
            return
        }
        if pPaths != nil && qPaths != nil {
            return
        }
        
        // logic
        paths = append(paths, r)
        if r.Val == p.Val {
            newL := make([]*TreeNode, len(paths))
            copy(newL, paths)
            pPaths = newL
        }
        
        if r.Val == q.Val {
            newL := make([]*TreeNode, len(paths))
            copy(newL, paths)
            qPaths = newL
        }
        dfs(r.Left, paths)
        dfs(r.Right, paths)
        paths = paths[:len(paths)-1]
    }
    dfs(root, nil)
    var out *TreeNode
    i := 0
    for i < len(pPaths) && i < len(qPaths) {
        if pPaths[i] == qPaths[i] {
            out = pPaths[i]
        }
        i++
    }
    return out
}