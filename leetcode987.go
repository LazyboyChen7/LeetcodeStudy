type st [][2]int
func verticalTraversal(root *TreeNode) [][]int {
    re := make([][]int, 0)
    m := make(map[int]st)
    dfs(m, root, 0, 0)
    var idx int = 1
    for i,_ := range m {
        if i < idx {
            idx = i
        }
    }
    for i:=0;i<len(m);i++ {
        s := m[idx+i]
        sort.Sort(s)
        res := make([]int,0)
        for _,v := range s {
            res = append(res, v[1])
        }
        re = append(re, res)
    }
    return re
}

func dfs(m map[int]st, node *TreeNode, x int, y int) {
    if node == nil {
        return
    }
    if _,ok := m[x]; !ok {
        m[x] = make(st,0)
    }
    m[x] = append(m[x], [2]int{y,node.Val})
    dfs(m, node.Left, x-1, y+1)
    dfs(m, node.Right, x+1, y+1)
}
func (s st)Len() int {
    return len(s)
}
func (s st)Less(i,j int) bool {
    if s[i][0] == s[j][0] {
        return s[i][1] < s[j][1]
    }
    return s[i][0] < s[j][0]
}
func (s st)Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}