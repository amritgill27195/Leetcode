/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sortedArrayToBST(nums []int) *TreeNode {
//     var dfs func(n []int) *TreeNode
//     dfs = func(n []int) *TreeNode {
//         // base
//         if n == nil || len(n) == 0 {return nil}
        
//         // logic
//         rootIdx := len(n) / 2
//         root := &TreeNode{Val: n[rootIdx]}
//         root.Left = dfs(n[:rootIdx])
//         root.Right = dfs(n[rootIdx+1:])
//         return root
//     }
    
//     return dfs(nums)
// }

func sortedArrayToBST(nums []int) *TreeNode {
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        // logic

        rootIdx := left + (right-left)/2 
        root := &TreeNode{Val: nums[rootIdx]}
        root.Left = dfs(left, rootIdx-1)
        root.Right = dfs(rootIdx+1, right)
        return root
    }
    
    return dfs(0, len(nums)-1)
}