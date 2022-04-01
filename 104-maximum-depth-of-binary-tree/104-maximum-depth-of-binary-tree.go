/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    Dfs: 2 ways
    1. maintain a height as your going down a branch
    2. base case will return the height in hand
    3. max(left, right) branch when recursion is going back up
    
    1. go down all the way first
    2. Then we hit bounce back from a nil branch ( it will return 0 )
    3. Then we will return max(left,right)+1  -- the +1 to count current node as part of height to return to parent node.
    time: o(n)
    space: o(h) - for the recursion stack, where h is the max height of tree
    
*/


// approach 1
func maxDepth(root *TreeNode) int {
    
    var dfs func(r *TreeNode, h int) int
    dfs = func (r *TreeNode, h int) int {
        if r == nil {
            return h
        }
        left := dfs(r.Left, h+1)
        right := dfs(r.Right, h+1)
        return int(math.Max(float64(left), float64(right)))
    }
    
    return dfs(root, 0)
}


// approach 2
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     left := maxDepth(root.Left)
//     right := maxDepth(root.Right)
//     return int(math.Max( float64(left), float64(right) ))+1
// }