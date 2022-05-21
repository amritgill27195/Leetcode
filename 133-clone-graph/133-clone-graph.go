/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {return nil}
    newNode := &Node{Val: node.Val}
    q := []*Node{node}
    visited := map[*Node]*Node{node: newNode}    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cloneOfDq := visited[dq]
        for _, n := range dq.Neighbors {
            if visited[n] == nil {
                clone := &Node{Val: n.Val}
                visited[n] = clone
                q = append(q, n)
            }
            cloneOfDq.Neighbors = append(cloneOfDq.Neighbors, visited[n])
        }
        
    }
    
    return newNode
    
}