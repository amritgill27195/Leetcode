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
    - Find paths that lead to p and q
    - Find the last common node between the 2 paths
    - The last common node is the lowest common ancestor
    
    time: o(n)
        - worse case our recursion goes left first and p and q are located in the right subtree as leaf nodes
        - plus we never exit early in 
    space: o(p+q)
        - o(h) for recursion stack + o(p+q) the paths array for both p and q which would be larger compared to o(h)
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var pPaths []*TreeNode
	var qPaths []*TreeNode

	var backtrack func(r *TreeNode, paths []*TreeNode)
	backtrack = func(r *TreeNode, paths []*TreeNode) {
		// base
        if r == nil{
			return
		}
        
        // exit early
        if len(pPaths) != 0 && len(qPaths) != 0 {return }
		// logic
		// action
        paths = append(paths, r)
        if r.Val == p.Val {
            // new list since we backtrack paths arr so that we dont end up with empty list ( since we remove last element from paths )
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