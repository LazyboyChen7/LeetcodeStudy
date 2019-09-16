var bcj []int
func findCircleNum(M [][]int) int {
    bcj = make([]int, len(M))
    for i:=range bcj {
        bcj[i] = i
    }
    for i:=0;i<len(M);i++ {
        for j:=0;j<len(M[0]);j++ {
            if M[i][j] == 1 {
                a := find(i)
                b := find(j)
                bcj[a] = b
            }
        }
    }
    re := 0
    for i:=range bcj {
        if bcj[i] == i {
            re++
        }
    }
    return re
}

func find(x int) int {
    if bcj[x] != x {
        bcj[x] = find(bcj[x])
    }
    return bcj[x]
}
