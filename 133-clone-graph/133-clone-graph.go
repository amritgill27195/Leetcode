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
    var dfs func(a *Node)
    dfs = func(a *Node) {
        // base
        if a == nil {return }
        if _, ok := visited[a]; ok {return} 
        
        // logic
        clone := &Node{Val:a.Val}
        visited[a] = clone
        for _, nh := range a.Neighbors {
            dfs(nh)
            clone.Neighbors = append(clone.Neighbors, visited[nh])
        }
    }
    dfs(node)
    return visited[node]
}