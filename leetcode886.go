func possibleBipartition(N int, dislikes [][]int) bool {
    arr := make([][]int, N)
    for i,_ := range arr {
        arr[i] = make([]int, 0)
    }
    color := make([]int, N)
    for _,v := range dislikes {
        arr[v[0]-1] = append(arr[v[0]-1], v[1]-1)
        arr[v[1]-1] = append(arr[v[1]-1], v[0]-1)
    }
    for i:=0;i<N;i++ {
        if color[i] == 0 && !dfs(i,arr,color, 1) {
            return false
        }
    }
    return true
}

func dfs(i int, arr [][]int, color []int, n int) bool {
    color[i] = n
    for _,v := range arr[i] {
        if color[v] == n {
            return false
        }
        if color[v] == 0 && !dfs(v,arr,color,-n) {
            return false
        }
    }
    return true
}