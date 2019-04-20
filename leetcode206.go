/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    add := &ListNode{-1,nil}
    add.Next = head
    re := head.Next
    pre := head
    for re != nil {
        pre.Next = re.Next
        re.Next = add.Next
        add.Next = re
        re = pre.Next
    }
    return add.Next
}