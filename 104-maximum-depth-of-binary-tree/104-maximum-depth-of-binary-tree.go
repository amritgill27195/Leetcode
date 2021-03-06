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
    
    BFS is also possible,
    level == height
    if we proccess a level -> height++
    
*/

// dfs approach
// time: o(n) - we visit all nodes
// space: o(h) - max height of the tree will be in our recursion stack
// func maxDepth(root *TreeNode) int {
//     var dfs func(r *TreeNode) int
//     dfs = func(r *TreeNode) int {
//         // base
//         if r == nil {
//             return 0
//         }
        
//         left := dfs(r.Left)
//         right := dfs(r.Right)
//         return int(math.Max(float64(left), float64(right)))+1
//     }
//     return dfs(root)
// }

// bfs approach
// time: o(n), we process all nodes level by level
// space: o(maxWidthOfTreeTakenByQueue)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q := []*TreeNode{root}
    h := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        h++
    }
    return h
}