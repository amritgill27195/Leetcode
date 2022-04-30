/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    cHead := &Node{Val: 0}
    cCurr := cHead
    
    origToDeepCopyNode := map[*Node]*Node{}
    
    current := head
    for current != nil {
        newNode := &Node{Val: current.Val}
        cHead.Next = newNode
        cHead = cHead.Next
        origToDeepCopyNode[current] = newNode
        current = current.Next
    }
    
    current = head
    for current != nil {
        if current.Random != nil {
            deepCpSrcNode := origToDeepCopyNode[current]
            deepCpDestNode := origToDeepCopyNode[current.Random]
            
            deepCpSrcNode.Random = deepCpDestNode
        }
        current = current.Next
    }
    return cCurr.Next
    
}