/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    return proc(root, 100000, -100000)
}

func proc(node *TreeNode, minV, maxV int) int {
    if node == nil {
        return 0
    }
    ans := max((node.Val - minV), (maxV - node.Val))
    if node.Val > maxV {
        maxV = node.Val
    }
    if node.Val < minV {
        minV = node.Val
    }
    ans = max(ans, max(proc(node.Left, minV, maxV), proc(node.Right, minV, maxV)))
    return ans
}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a 
}