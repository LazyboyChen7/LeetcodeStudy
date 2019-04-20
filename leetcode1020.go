func numEnclaves(A [][]int) int {
    rowlen := len(A)
    collen := len(A[0])
    for i:=0 ;i<collen; i++ {
        if A[0][i] == 1{
            proc(A, 0, i)
        }
        if A[rowlen-1][i] == 1 {
            proc(A, rowlen-1, i)
        }
    }
    for i:=0 ;i<rowlen; i++ {
        if A[i][0] == 1 {
            proc(A, i, 0)
        }
        if A[i][collen-1] == 1 {
            proc(A, i, collen-1)    
        }
    }
    var re int
    for i:=0;i<rowlen;i++ {
        for j:=0;j<collen;j++ {
            if A[i][j] == 1 {
                re++
            }
        }
    }
    return re
}

func proc(A [][]int, i,j int) {
    if A[i][j] == 1 {
        A[i][j] = 0
        if i>0 {
            proc(A, i-1, j)
        }
        if i<len(A)-1 {
            proc(A, i+1, j)
        }
        if j>0 {
            proc(A, i, j-1)
        }
        if j<len(A[0])-1 {
            proc(A, i, j+1)
        }
    }
}