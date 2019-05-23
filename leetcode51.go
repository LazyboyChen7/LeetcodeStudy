var re [][]string
func solveNQueens(n int) [][]string {
    re = make([][]string, 0)
    board := make([][]bool, n)
    for i,_ := range board {
        board[i] = make([]bool, n)
    }
    dfs(n, board)
    return re
}

func dfs(dep int, b [][]bool) {
    if dep == 0 {
        proc(b)
        return
    }
    r := len(b)-dep
    for c:=0;c<len(b);c++ {
        if !isValid(b, r, c) {
            continue
        }
        b[r][c] = true
        dfs(dep-1, b)
        b[r][c] = false
    }
}

func isValid(b [][]bool, r,c int) bool {
    for i:=0;i<=r;i++ {
        if b[i][c] {
            return false
        }
        if c>=i && b[r-i][c-i] {
            return false
        }
        if c+i<len(b) && b[r-i][c+i] {
            return false
        }
    }
    return true
}

func proc(b [][]bool) {
    ss := make([]string, len(b))
    for i:=0;i<len(b);i++ {
        c := make([]byte, len(b))
        for j:=0;j<len(b);j++ {
            if b[i][j] {
                c[j] = 'Q'
            } else {
                c[j] = '.'
            }
        }
        ss[i] = string(c)
    }
    re = append(re, ss)
}
