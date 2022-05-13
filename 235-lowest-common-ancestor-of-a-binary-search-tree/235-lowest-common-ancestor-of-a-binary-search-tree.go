/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for true {
        if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else if q.Val < root.Val && p.Val < root.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return root
}