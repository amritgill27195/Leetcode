/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    s := []*TreeNode{}
    for root != nil || len(s) != 0 {
        for root != nil {
            s = append(s, root)
            root = root.Left
        }
        root = s[len(s)-1]
        s = s[:len(s)-1]
        k--
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
    return -1
}