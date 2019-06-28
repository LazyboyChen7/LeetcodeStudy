/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || n==m {
        return head
    }
    pre := &ListNode{0,head}
    re := pre
    for m>1 {
        re = re.Next
        m--
        n--
    }
    d := re
    frt := &ListNode{0,nil}
    for n > 0 {
        re = re.Next
        newNode := &ListNode{re.Val, frt.Next}
        frt.Next = newNode
        n--
    }
    d.Next = frt.Next
    for d.Next != nil {
        d = d.Next
    }
    d.Next = re.Next
    return pre.Next
}