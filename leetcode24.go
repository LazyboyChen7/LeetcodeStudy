/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p,q,k := head, head.Next, &ListNode{0,head}
    for q != nil {
        if k.Next == head {
            head = k
        }
        k.Next = q
        p.Next = q.Next
        q.Next = p
        k = p
        p = p.Next
        if p == nil {
            return head.Next
        }
        q = p.Next
    }
    return head.Next
}