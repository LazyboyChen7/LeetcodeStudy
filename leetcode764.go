func orderOfLargestPlusSign(N int, mines [][]int) int {
    arr := make([][]int, N)
    for i,_ := range arr {
        arr[i] = make([]int, N)
        for j,_ := range arr[i] {
            arr[i][j] = N
        }
    }
    for _,v := range mines {
        arr[v[0]][v[1]] = 0
    }
    for i:=0;i<N;i++ {
        var l,r,u,d int
        for j,k:=0,N-1;j<N;j,k=j+1,k-1 {
            l = proc(arr[i][j], l)
            arr[i][j] = min(arr[i][j], l)
            r = proc(arr[i][k], r)
            arr[i][k] = min(arr[i][k], r)
            u = proc(arr[j][i], u)
            arr[j][i] = min(arr[j][i], u)
            d = proc(arr[k][i], d)
            arr[k][i] = min(arr[k][i], d)
        }
    }
    var re int
    for _,v := range arr {
        for _,n := range v {
            re = max(re, n)
        }
    }
    return re
}

func proc(a,b int) int {
    if a == 0 {
        return 0
    }
    return b+1
}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}
