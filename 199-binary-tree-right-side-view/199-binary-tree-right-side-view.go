/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {return nil}
    out := []int{}
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q)
        tmp := len(q)
        for qSize != 0 {
            if tmp == qSize {
                out = append(out, q[qSize-1].Val)
            }
            dq := q[0]
            q = q[1:]
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
    }
    return out
}