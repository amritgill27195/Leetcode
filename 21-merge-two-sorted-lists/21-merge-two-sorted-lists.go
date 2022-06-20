/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {return nil}
    if list1 != nil && list2 == nil {return list1}
    if list1 == nil && list2 != nil {return list2}
    
    out := &ListNode{Val: 0}
    tail := out
    
    for list1 != nil || list2 != nil {
        
        var l1Val int = -101
        if list1 != nil {
            l1Val = list1.Val
        }
        
        var l2Val int = -101
        if list2 != nil {
            l2Val = list2.Val
        }
        
        
        if l1Val != -101 && l2Val != -101 {
            if l1Val < l2Val {
                tail.Next = &ListNode{Val: l1Val}
                list1 = list1.Next
            } else {
                tail.Next = &ListNode{Val: l2Val}
                list2 = list2.Next
            }
        } else if l1Val != -101 {
            tail.Next = &ListNode{Val: l1Val}
            list1 = list1.Next
        } else if l2Val != -101 {
            tail.Next = &ListNode{Val: l2Val}
            list2 = list2.Next
        }
        tail = tail.Next
    }
    return out.Next
}