/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	slc := make([]*TreeNode, 0)
	arr := proc(S)
	root := &TreeNode{}
	for i := 0; i < len(S) && S[i] != '-'; {
		root.Val = root.Val*10 + int(S[i]) - '0'
		i++
	}
	slc = append(slc, root)
	for i := 2; i < len(arr); i += 2 {
		if arr[i] == len(slc) {
			node := &TreeNode{
				Val: arr[i+1],
			}
			slc[len(slc)-1].Left = node
			slc = append(slc, node)
		} else {
			for arr[i] != len(slc) {
				slc = slc[:len(slc)-1]
			}
			node := &TreeNode{
				Val: arr[i+1],
			}
			slc[len(slc)-1].Right = node
			slc = append(slc, node)
		}
	}
	return root
}

func proc(str string) []int {
	arr := make([]int, 0)
	s := []byte(str)
	for i := 0; i < len(s); {
		var n int
		for i < len(s) && s[i] == '-' {
			n++
			i++
		}
		arr = append(arr, n)
		var num int
		for i < len(s) && s[i] != '-' {
			num = num*10 + int(s[i]) - '0'
			i++
			//fmt.Println(num)
		}
		arr = append(arr, num)
	}
	return arr
}