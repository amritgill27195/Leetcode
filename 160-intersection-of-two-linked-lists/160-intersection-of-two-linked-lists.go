/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    /*
        approach 1 : using hashmap
        - loop over headA and headB
        - attempt to add both node refs in hashmap
        - if map already contains a nodeRef, that means that node is the intersection point
    */
    if headA == nil || headB == nil {
        return nil
    }
    
    nodeRefs := map[*ListNode]struct{}{}
    for headA != nil || headB != nil {
        if headA != nil {
            _, ok := nodeRefs[headA]
            if ok {
                return headA
            }
            nodeRefs[headA] = struct{}{}
            headA = headA.Next
        }
        
        if headB != nil {
            _, ok := nodeRefs[headB]
            if ok {
                return headB
            }
            nodeRefs[headB] = struct{}{}
            headB = headB.Next
        }
    }
    
    return nil
}