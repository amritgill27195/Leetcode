/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
//     var dfs func(min, max int, r *TreeNode) bool
//     dfs = func(min, max int, r *TreeNode) bool {
//         // base
//         if r == nil {return true}
//         // logic
//         if r.Val <= min || r.Val >= max {
//             return false
//         }
//         return dfs(min, r.Val, r.Left) && dfs(r.Val, max, r.Right)
//     }
//     return dfs(math.MinInt64, math.MaxInt64, root)
// }


func isValidBST(root *TreeNode) bool {
    type sNode struct {
        n *TreeNode
        min int
        max int
    }
    
    // [ [*TreeNode, min, max] ]
    s := []*sNode{ &sNode{n:root, min: math.MinInt64, max: math.MaxInt64} }

    for len(s) != 0 {
        top := s[len(s)-1]
        s= s[:len(s)-1]
        
        node := top.n
        min := top.min
        max := top.max
        
        if node.Val <= min || node.Val >= max {
            return false
        }
        
        if node.Left != nil { s = append(s, &sNode{n: node.Left, min: min, max: node.Val})}
        if node.Right != nil { s = append(s, &sNode{n: node.Right, min: node.Val, max: max})}
        
    }
    return true
}