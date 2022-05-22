/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    curr := head
    for curr != nil {
        next := curr.Next
        deepCp := &Node{Val: curr.Val}
        curr.Next = deepCp
        deepCp.Next= next
        curr = next
    }
    
    curr = head
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = next
    }
    
    out := &Node{Val: 0}
    tail := out
    curr = head
    // 1-2-2`-3-3`
    // c 
    //   t
    for curr != nil {
        next := curr.Next.Next
        tail.Next = curr.Next
        tail = tail.Next
        curr.Next = next
        curr = next
    }
    return out.Next
}