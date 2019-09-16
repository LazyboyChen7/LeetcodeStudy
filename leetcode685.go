var parent []int
var bcj []int
func findRedundantDirectedConnection(edges [][]int) []int {
    parent,bcj = make([]int,1001),make([]int,1001)
    re,edge1,edge2 := []int{},[]int{},[]int{}
    for i:=range parent {
        bcj[i] = i
    }
    for _,v := range edges {
        if parent[v[1]] != 0 {
            edge1 = []int{parent[v[1]],v[1]}
            edge2 = v
        } else {
            parent[v[1]] = v[0]
            v1,v2 := find(v[0]),find(v[1])
            if v1 == v2 {
                re = v
            } else {
                bcj[v1] = v2
            }
        }
    }
    if len(edge1) != 0 && len(edge2) != 0 {
        if len(re) == 0 {
            return edge2
        }
        return edge1
    }
    return re
}

func find(x int) int {
    if bcj[x] != x {
        bcj[x] = find(bcj[x])
    }
    return bcj[x]
}
