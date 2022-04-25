/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
/*
    approach: level order using BFS
    - At each level,
    - For each node we dq, we will set the dq.Next to top queue.peek 
        - Unless its the last node of the level ( farthest right )
    - Which means we will take a queue size
    
    time: o(n)
    space: o(maxWidthOfTree)
*/
// func connect(root *Node) *Node {
//     if root == nil {
//         return nil
//     }
//     q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for i := 0; i < qSize; i++ {
//             dq := q[0]
//             q = q[1:]
//             if i < qSize-1 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//         }
//     }
//     return root
// }


/*
    approach
*/

func connect(root *Node) *Node {
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {
            return
        }
        
        // logic
        if r.Left != nil && r.Right != nil {
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