/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// just traverse with bst constraints 
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     b := &bst{sum: 0}
//     b.inorderDfs(root, low, high)
//     return b.sum
// }

// type bst struct {
//     sum int
// }
// func (b *bst)inorderDfs(root *TreeNode, low, high int) {
//     // base 
//     if root == nil{
//         return
//     }
    
//     if low < root.Val && root.Val < high {
//         b.inorderDfs(root.Left, low, high)
//     }    
    
//     if root.Val >= low && root.Val <= high {
//         b.sum += root.Val
//     }
    
//     if low < root.Val && root.Val < high{ 
//         b.inorderDfs(root.Right, low, high)
//     }
// }


// plain postOrder traversal
// time: o(n)
// space: o(h)
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     sum := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
       
//         // logic
//         dfs(r.Left)
//         dfs(r.Right)
//         if r.Val >= low && r.Val <= high {
//             sum+=r.Val
//         }
//     }
//     dfs(root)
//     return sum
// }

// plain preorder traversal
// time: o(n)
// space: o(h)
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     sum := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
       
//         if r.Val >= low && r.Val <= high {
//             sum+=r.Val
//         }
//         // logic
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return sum
// }


// plain inorder traversal
// time: o(n)
// space: o(h)
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        // logic
        dfs(r.Left)
        if r.Val >= low && r.Val <= high {
            sum+=r.Val
        }

        dfs(r.Right)
    }
    dfs(root)
    return sum
}