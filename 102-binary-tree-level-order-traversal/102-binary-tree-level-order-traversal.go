/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iterative using BFS
// time: o(n)
// space: o(maxWidthOfTheTree)
// func levelOrder(root *TreeNode) [][]int {
//     if root == nil {
//         return nil
//     }
//     out := [][]int{}
//     queue := []*TreeNode{root}
    
//     for len(queue) != 0 {
//         qSize := len(queue)
//         cl := []int{}
//         for qSize != 0 {
//             dq := queue[0]
//             queue = queue[1:]
//             cl = append(cl, dq.Val)
//             if dq.Left != nil {queue = append(queue, dq.Left)}
//             if dq.Right != nil {queue = append(queue, dq.Right)}

//             qSize--
//         }
//         out = append(out, cl)
//     }
//     return out
// }



// recursive using DFS
// time: o(n)
// space: o(h) -- at worse, we will have the max height number of nodes in our recursion stack

type sol struct {
    result [][]int
}

func levelOrder(root *TreeNode) [][]int{
    sol := new(sol)
    sol.result = [][]int{}
    sol.preorderDfs(root, 0)
    return sol.result
}


func (s *sol) preorderDfs(root *TreeNode, level int) {
    
    // base
    if root == nil {
        return
    }
    
    //logic
    if len(s.result) == level {
        s.result = append(s.result, []int{})
    }
    s.result[level] = append(s.result[level], root.Val)
    level = level+1
    
    s.preorderDfs(root.Left, level)
    s.preorderDfs(root.Right, level)
}
