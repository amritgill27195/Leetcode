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



// approach : level order using DFS ( by maintaining a level at each node in recursion stack )
// level order can be done using DFS as well
// At each node in the recursion stack we will maintain a level -- which represents the idx in the nested array
// if the len of the result array == level - it means result[level] idx does not exist - for it to exist, the len > level.
// therefore we will initialize a new array within, so that result[level] exists.
// once the result[level] array does exist, we will simply append the current level root.Val to result[level]

// time: o(n)
// space: o(h) -- at worse, we will have the max height number of nodes in our recursion stack

// tried passing result as a reference but golang complains that *result[idx] indexing on references are not supported
// so then the class solution to narrow down the global scope of result to only 1 instance or else global pollution happens
type sol struct {
    result [][]int
}

func levelOrder(root *TreeNode) [][]int{
    sol := new(sol)
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
