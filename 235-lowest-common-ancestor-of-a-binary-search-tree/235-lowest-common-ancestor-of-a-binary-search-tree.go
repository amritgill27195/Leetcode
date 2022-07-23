/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if root.Val == p.Val || root.Val == q.Val {
            return root
        }
        if p.Val > root.Val && q.Val > root.Val {
           root = root.Right
            continue
        }
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
            continue
        }
        break
    }
    return root
}