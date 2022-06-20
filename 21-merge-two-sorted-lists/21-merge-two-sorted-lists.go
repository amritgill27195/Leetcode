/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
2 pointers
1. Start at l1
2. Start at l2
3. While both are not nil,
    a. Compare which is smaller and append to output tail
    b. the smaller pointer moves forward
4. Finally append whoever did not hit nil ( in case of uneven sized LL )

time: o(l1+l2) 
space: o(1)

*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    l1 := list1
    l2 := list2
    out := &ListNode{Val: 0}
    tail := out
    for l1 != nil && l2 != nil {
        l1Val := l1.Val
        l2Val := l2.Val
        if l1Val < l2Val {
            tail.Next = &ListNode{Val: l1Val}
            tail = tail.Next
            l1 = l1.Next
        } else {
            tail.Next = &ListNode{Val: l2Val}
            tail = tail.Next
            l2 = l2.Next
        }
    }
    for l1 != nil {
        tail.Next = &ListNode{Val: l1.Val}
        tail = tail.Next
        l1 = l1.Next
    }
    for l2 != nil {
        tail.Next = &ListNode{Val: l2.Val}
        tail = tail.Next
        l2 = l2.Next
    }
    return out.Next
}