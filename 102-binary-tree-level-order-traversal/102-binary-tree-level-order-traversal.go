/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    out := [][]int{}
    queue := []*TreeNode{root}
    
    for len(queue) != 0 {
        qSize := len(queue)
        cl := []int{}
        for qSize != 0 {
            dq := queue[0]
            queue = queue[1:]
            cl = append(cl, dq.Val)
            if dq.Left != nil {queue = append(queue, dq.Left)}
            if dq.Right != nil {queue = append(queue, dq.Right)}

            qSize--
        }
        out = append(out, cl)
    }
    return out
}

