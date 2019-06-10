func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if !proc(root, 0, limit) {
        return nil
    }
    return root
}

func proc(node *TreeNode, sum int, limit int) (bool) {
    if node.Left == nil && node.Right == nil {
        sum += node.Val
        if sum < limit {
            return false
        } else {
            return true
        }
    }
    f1,f2 := true,true
    if node.Left != nil {
        f1 = !proc(node.Left, sum+node.Val, limit)
    }
    if node.Right != nil {
        f2 = !proc(node.Right, sum+node.Val, limit)
    }
    if f1 && f2 {
        node = nil
        return false
    } else if f1 {
        node.Left = nil
    } else if f2 {
        node.Right = nil
    }
    return true
}
