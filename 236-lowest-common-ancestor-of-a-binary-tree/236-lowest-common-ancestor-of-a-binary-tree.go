/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: brute force
    - Find paths that lead to p
    - Find paths that lead to q
    - Find the last common node between the 2 paths
    - The last common node is the lowest common ancestor
    
    time: o(n)
        - worse case our recursion goes left first and p and q are located in the right subtree as leaf nodes
    space: 
        - o(h) to find p + o(h) to find q + o(h) path to p + o(h) path to q
        - o(h)
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var pPaths []*TreeNode
	var qPaths []*TreeNode

	var backtrack func(r *TreeNode, paths []*TreeNode)
	backtrack = func(r *TreeNode, paths []*TreeNode) {
		// base
		if r == nil {
			return
		}

		// logic
		// action
        paths = append(paths, r)
        if r.Val == p.Val {
			newP := make([]*TreeNode, len(paths))
			copy(newP, paths)
			pPaths = newP
		}
		if r.Val == q.Val {
			newQ := make([]*TreeNode, len(paths))
			copy(newQ, paths)
			qPaths = newQ
		}
		// recurse
		backtrack(r.Left, paths)
		backtrack(r.Right, paths)
		// backtrack
		paths = paths[:len(paths)-1]
	}
	backtrack(root, nil)

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


/*
    approach:
    time:
    space:
*/
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    
// }