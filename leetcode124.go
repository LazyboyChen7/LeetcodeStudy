func maxPathSum(root *TreeNode) int {
    var re int = ^int(^uint(0) >> 1)
    proc(root, &re)
    return re
}

func proc(node *TreeNode, re *int) int {
    if node == nil {
        return 0
    }
    left := max(proc(node.Left, re), 0)
    right := max(proc(node.Right, re), 0)
    *re = max(*re, node.Val + left + right)
    return node.Val + max(left, right)
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}