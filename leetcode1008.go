// 1008. Construct Binary Search Tree from Preorder Traversal
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) < 1 {
		return nil
	}
	node := new(TreeNode)
	node.Val = preorder[0]
	l := 1
	for l < len(preorder) && preorder[l] < node.Val {
		l++
	}
	node.Left = bstFromPreorder(preorder[1:l])
	node.Right = bstFromPreorder(preorder[l:])
	return node
}

// 执行用时 : 4 ms, 在Construct Binary Search Tree from Preorder Traversal的Go提交中击败了100.00% 的用户
// 内存消耗 : 3.2 MB, 在Construct Binary Search Tree from Preorder Traversal的Go提交中击败了100.00% 的用户