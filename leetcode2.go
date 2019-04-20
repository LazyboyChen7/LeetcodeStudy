func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{-1,nil}
    h:= head
    for count:=0;l1!=nil||l2!=nil||count!=0;{
        if l1!=nil{
            count+= l1.Val      
            l1 = l1.Next
        }
        if l2!=nil{
            count+= l2.Val
            l2 = l2.Next
        }
        node:= &ListNode{count%10,nil}
        head.Next = node
        head = head.Next
        count = count/10
    }
    return h.Next
}