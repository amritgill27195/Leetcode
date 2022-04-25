/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    BST == running inorder gives us a sorted output  ( as long as its a valid bst )
    approach: brute force
    - run inorder
    - at each node, store the node val in an array
        - this array will be global and cannot be passed as arg because reference data structures combined with pass-by-value slice abstraction in golang will output incorrect answers
    - finally, return the kth smallest item arry from inorder ( idx: k-1 )

*/

func kthSmallest(root *TreeNode, k int) int {
    result := []int{}
    var inorder func(r *TreeNode)
    inorder = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        inorder(r.Left)
        result = append(result, r.Val)
        inorder(r.Right)
    }
    inorder(root)
    return result[k-1]
}
// dfs approach
// iterative inorder approach
// time: o(k)
// space: o(h)
// func kthSmallest(root *TreeNode, k int) int {
//     var inorder func(r *TreeNode) int
//     inorder = func(r *TreeNode) int {
//         // base
//         if r == nil {
//             return -1
//         }
        
        
//         // logic
//         left := inorder(r.Left)
//         if left != -1 {
//             return left
//         }
//         k--
//         if k == 0 {
//             return r.Val
//         }
//         right := inorder(r.Right)
//         if right != -1 {
//             return right
//         }
//         return -1
//     }
//     return inorder(root)
// }


// iterative inorder approach
// time: o(k)
// space: o(1)
// func kthSmallest(root *TreeNode, k int) int {
//     s := []*TreeNode{}
//     for root != nil || len(s) != 0 {
//         for root != nil {
//             s = append(s, root)
//             root = root.Left
//         }
//         root = s[len(s)-1]
//         s = s[:len(s)-1]
//         k--
//         if k == 0 {
//             return root.Val
//         }
//         root = root.Right
//     }
//     return -1
// }