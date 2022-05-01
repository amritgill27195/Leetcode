/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */


// classis bfs level order using a queue
// func connect(root *Node) *Node {
//     if root == nil {return root}	
//     q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for i := 0; i < qSize; i++ {
//             dq := q[0]
//             q = q[1:]
//             if i != qSize-1 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {
//                 q = append(q, dq.Left)
//             }
//             if dq.Right != nil {
//                 q = append(q, dq.Right)
//             }
//         }
//     }
//     return root 
// }

// DFS
func connect(root *Node) *Node {
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {
            return
        }
        // logic
        if r.Left != nil {
            r.Left.Next = r.Right
        }
        
        if r.Next != nil && r.Right != nil {
            r.Right.Next = r.Next.Left
        }
    
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return root
}