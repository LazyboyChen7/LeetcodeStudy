/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    p,q := head,head.Next
    for q != nil && q.Next != nil{
        p = p.Next
        q = q.Next.Next
    }
    h2 := p.Next
    p.Next = nil
    h2 = reverse(h2)
    re := head
    for h2 != nil {
        p := h2
        h2 = h2.Next
        p.Next = re.Next
        re.Next = p
        re = p.Next
    }
}

func reverse(h *ListNode) *ListNode {
    p := &ListNode{0,h}
    n := h.Next
    for n != nil {
        h.Next = n.Next
        n.Next = p.Next
        p.Next = n
        n = h.Next
    }
    return p.Next
}