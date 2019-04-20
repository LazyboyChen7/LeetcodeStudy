func nextLargerNodes(head *ListNode) []int {
    arr := make([]int, 0)
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
	re := make([]int, 0)
	stk := make([]int, 0)
    for i := 0; i<len(arr); i++ {
		re = append(re, -1)
		for len(stk) != 0 {
			if arr[i] > arr[stk[len(stk)-1]] {
				re[stk[len(stk)-1]] = arr[i]
				stk = stk[:len(stk)-1]
			} else {
				break
			}
		}
		stk = append(stk, i)
	}
	for i := 0; i < len(stk); i++ {
		re[stk[i]] = 0
	}
	return re
}