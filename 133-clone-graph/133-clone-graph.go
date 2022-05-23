/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
// BFS
func cloneGraph(node *Node) *Node {
    if node == nil {return nil}
    visited := map[*Node]*Node{}
    clone := &Node{Val: node.Val}
    visited[node] = clone
    q := []*Node{node}
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        
        for _, nh := range dq.Neighbors {
            _, ok := visited[nh]
            if !ok {
                visited[nh] = &Node{Val: nh.Val}
                q = append(q, nh)
            }
            visited[dq].Neighbors = append(visited[dq].Neighbors, visited[nh])
        }
    }
    return clone
}