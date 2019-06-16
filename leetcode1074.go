var re int
func numSubmatrixSumTarget(matrix [][]int, target int) int {
    re = 0
    arr := proc(matrix,target)
    for i:=0;i<len(matrix);i++ {
        for j:=0;j<len(matrix[0]);j++ {
            com(i,j,arr,matrix,target)
        }
    }
    return re
}

func com(a,b int, arr,slc [][]int, t int) {
    for i:=a+1;i<len(arr);i++ {
        for j:=b+1;j<len(arr[0]);j++ {
            if arr[i][j] - arr[i][b] -arr[a][j] + arr[a][b] == t {
                re++
            }
        }
    }
}


func proc(slc [][]int, t int) [][]int{
    arr := make([][]int, len(slc)+1)
    for i,_ := range arr {
        arr[i] = make([]int, len(slc[0])+1)
    }
    for i:=1;i<len(arr);i++ {
        for j:=1;j<len(arr[0]);j++ {
            arr[i][j] = arr[i-1][j] + arr[i][j-1] - arr[i-1][j-1] + slc[i-1][j-1]
        }
    }
    return arr
}
