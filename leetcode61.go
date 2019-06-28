/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    test, length := head, 0
    for test != nil {
        test = test.Next
        length++
    }
    k = k%length
    if k == 0 {
        return head
    }
    p,q,f := head, head, &ListNode{0,head}
    for k != 1 {
        q = q.Next
        k--
    }
    for q.Next != nil {
        p = p.Next
        q = q.Next
        f = f.Next
    }
    f.Next = nil
    q.Next = head
    return p
}