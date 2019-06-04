var dx = []int{-1,-1,-1,0,0,0,1,1,1}
var dy = []int{-1,0,1,-1,0,1,-1,0,1}
var ib = []byte{'B','1','2','3','4','5','6','7','8'}

func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    }
    proc(board, click[0], click[1])
    return board
}

func proc(b [][]byte, r,c int) {
    var n int = search(b,r,c)
    b[r][c] = ib[n]
    for i:=0;i<9 && n==0;i++ {
        x,y := r+dx[i],c+dy[i]
        if x > -1 && x < len(b) && y > -1 && y < len(b[0]) && b[x][y] == 'E' {
            proc(b, x, y)
        }
    }
}

func search(b [][]byte, r,c int) int {
    var re int
    for i,_ := range dx {
        x,y := r+dx[i],c+dy[i]
        if x > -1 && x < len(b) && y > -1 && y < len(b[0]) && b[x][y] == 'M' {
            re++
        }
    }
    return re
}