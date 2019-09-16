var parent []int
func findRedundantConnection(edges [][]int) []int {
    parent = make([]int,1001)
    re := []int{}
    for i:=range parent {
        parent[i] = i
    }
    for _,v := range edges {
        v1,v2 := find(v[0]),find(v[1])
        if v1 == v2 {
            re = v
        } else {
            parent[v1] = v2
        }
    }
    return re
}

func find(x int) int {
    if parent[x] != x {
        parent[x] = find(parent[x])
    }
    return parent[x]
}
