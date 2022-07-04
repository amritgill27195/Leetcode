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
    
    // 1-1|-2-2|-3-3|-nil
    curr := head
    for curr != nil {
        next := curr.Next
        cp := &Node{Val: curr.Val}
        curr.Next = cp
        cp.Next = next
        curr = next
    }
    // now for each copy node, connect to its corresponding random ref
    curr = head
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = next
    }
    
    curr = head
    out := &Node{Val: 0}
    tail := out
    var prev *Node

    // 1-1|-2-2|-3-3|-nil
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        cp := curr.Next
        curr.Next = nil
        tail.Next = cp
        tail = tail.Next
        if prev != nil {
            prev.Next = curr
        }
        prev = curr
        curr = next
    }
    
    return out.Next
    
}