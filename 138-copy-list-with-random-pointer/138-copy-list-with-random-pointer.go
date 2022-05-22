/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
    srcToDeepCp := map[*Node]*Node{}
    out := &Node{Val: 0}
    tail := out
    
    current := head
    for current != nil {
        deepCp := &Node{Val: current.Val}
        tail.Next = deepCp
        tail = tail.Next
        srcToDeepCp[current] = deepCp
        current = current.Next
    }
    
    current = head
    for current != nil {
        if current.Random != nil {
            src := srcToDeepCp[current]
            dest := srcToDeepCp[current.Random]
            src.Random = dest
        }
        current = current.Next
    }
    
    return out.Next
    
}