func solveSudoku(board [][]byte)  {
    for proc(board) {}
    dfs(board,0,0)
}

func proc(s [][]byte) bool {
    var r bool
    for i:=0;i<9;i++ {
        for j:=0;j<9;j++ {
            if s[i][j] == '.' {
                set(s, i, j)
                if s[i][j] != '.' {
                    r = true
                }
            }
        }
    }
    return r
}

func set(s [][]byte, i,j int) {
    var n int
    for k:=0;k<9;k++ {
        if s[i][k] != '.' {
            n |= 1 << (s[i][k] - '0')   
        }
        if s[k][j] != '.' {
            n |= 1 << (s[k][j] - '0')
        }
    }
    for r:=i/3*3;r<i/3*3+3;r++ {
        for c:=j/3*3;c<j/3*3+3;c++ {
            if s[r][c] != '.' {
                n |= 1 << (s[r][c] - '0') 
            }
        }
    }
    num := 1<<10 - 2 - n
    if num&(num-1) == 0 {
        s[i][j] = byte(tsq(num)+'0')
    }
}

func tsq(n int) int {
    var re int
    for n != 1 {
        re++
        n >>= 1
    }
    return re
}

func dfs(board [][]byte, i,j int) bool {
    if i==9 {
        return true
    }
    if j>8 {
        return dfs(board, i+1, 0)
    }
    if board[i][j] == '.' {
        for k:=1;k<=9;k++ {
            board[i][j] = byte(k + '0')
            if isValid(board, i, j) {
                if dfs(board, i, j+1) {
                    return true
                }
            }
            board[i][j] = '.'
        }
    } else {
        return dfs(board, i, j+1)
    }
    return false
}

func isValid(b [][]byte, i,j int) bool {
    for k:=0;k<9;k++ {
        if k != i && b[i][j] == b[k][j] {
            return false
        }
        if k != j && b[i][j] == b[i][k] {
            return false
        }
    }
    for r := i/3*3;r<i/3*3+3;r++ {
        for c := j/3*3;c<j/3*3+3;c++ {
            if (r!=i || c!=j) && b[i][j] == b[r][c] {
                return false
            }
        }
    }
    return true
}