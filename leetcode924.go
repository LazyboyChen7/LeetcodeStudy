func minMalwareSpread(graph [][]int, initial []int) int {
    dist := make([]int, len(graph))
    size := make([]int, len(graph))
    for i:=0;i<len(graph);i++ {
        dist[i] = i
        size[i] = 1
    }
    for i:=0;i<len(graph);i++ {
        fi := find(dist,i)
        for j:=i+1;j<len(graph);j++ {
            if graph[i][j] == 1 {
                fj := find(dist, j)
                if fi != fj {
                    dist[fj] = fi
                    size[fi] += size[fj]
                }
            }
        }
    }
    m,re := 0,0
    for i:=0;i<len(initial);i++ {
        n := size[find(dist, initial[i])]
        if n > m {
            m = n
            re = initial[i]
        } else if n == m {
            re = min(re, initial[i])
        }
    }
    return re
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}

func find(s []int, n int) int {
    for n != s[n] {
        n = s[n]
    }
    return n
}