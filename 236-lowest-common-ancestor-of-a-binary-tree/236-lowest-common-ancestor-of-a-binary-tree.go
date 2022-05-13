/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    
    var dfs func(r *TreeNode, paths []*TreeNode, t *TreeNode) []*TreeNode
    dfs = func(r *TreeNode, paths []*TreeNode, t *TreeNode) []*TreeNode {
        // base
        if r == nil {
            return nil
        }
        // logic
        paths = append(paths, r)
        if r.Val == t.Val {
            newL := make([]*TreeNode, len(paths))
            copy(newL, paths)
            return newL
        }
        
        left := dfs(r.Left, paths, t)
        if left != nil {
            return left
        }
        return dfs(r.Right, paths, t)
    }
    
    var pPaths []*TreeNode = dfs(root, nil, p)
    var qPaths []*TreeNode = dfs(root, nil, q)
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