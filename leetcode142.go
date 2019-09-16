func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    s,f := head,head
    for f != nil && f.Next != nil {
        s = s.Next
        f = f.Next.Next
        if s == f {
            s = head
            for s != f {
                s = s.Next
                f = f.Next
            }
            return s
        }
    }
    return nil
}
