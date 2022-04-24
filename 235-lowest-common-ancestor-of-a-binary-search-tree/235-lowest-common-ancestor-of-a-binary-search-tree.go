/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(a *TreeNode, target *TreeNode, paths []*TreeNode) []*TreeNode
    dfs = func(a *TreeNode, target *TreeNode, paths []*TreeNode) []*TreeNode {
        // base
        if a == nil || target == nil {
            return nil
        }
        paths = append(paths, a)
        
        // logic
        if a.Val == target.Val {
            return paths
        }
        if target.Val < a.Val {
            return dfs(a.Left, target, paths)
        }
        return dfs(a.Right, target, paths)
    }
    pPaths := dfs(root, p, nil)
    qPaths := dfs(root, q, nil)
    
    var out *TreeNode
    pPtr := 0
    qPtr := 0
    for pPtr < len(pPaths) && qPtr < len(qPaths) {
        if pPaths[pPtr] == qPaths[qPtr] {
            out = pPaths[pPtr] 
        }
        pPtr++
        qPtr++
    }
    return out
}