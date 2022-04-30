/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time: o(n)
// space: o(h)
// approach recursion
// func flatten(root *TreeNode)  {
    
//     // base
//     if root == nil {
//         return
//     }
    
//     //logic
//     flatten(root.Left)
//     if root.Left != nil {
//         tmp := root.Right
//         root.Right = root.Left
//         root.Left = nil
//         // go right as much as possible to get to the tail end of the linkedlist
//         for root.Right != nil {
//             root = root.Right
//         }
//         // once we have reached the tail end, set the tail.Right to tmp ( prev right )
//         root.Right = tmp
//     }
//     flatten(root.Right)    
// }

// iterative
func flatten(root *TreeNode)  {
    var stack []*TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Left != nil {
            tmp := root.Right
            root.Right = root.Left
            root.Left = nil
            for root.Right != nil {
                root = root.Right
            }
            root.Right = tmp
        }
        root = root.Right
    }

}