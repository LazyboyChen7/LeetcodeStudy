var r,c int
type v [][]int
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    // arr := make([][]bool, R)
    // for i:=0;i<R;i++ {
    //     arr[i] = make([]bool, C)
    // }
    // fmt.Println(len(arr),r0,len(arr[0]),c0)
    // arr[r0][c0] = true
    // slc := make([][]int, 0)
    // re := make([][]int, 0)
    // slc = append(slc, []int{r0,c0})
    // re = append(re, []int{r0,c0})
    // for len(slc) != 0 {
    //     s := slc[0]
    //     slc = slc[1:]
    //     if s[0]<R-1 && !arr[s[0]+1][s[1]] {
    //         slc = append(slc, []int{s[0]+1, s[1]})
    //         re = append(re, []int{s[0]+1, s[1]})
    //         arr[s[0]+1][s[1]] = true
    //     }
    //     if s[0]>0 && !arr[s[0]-1][s[1]] {
    //         slc = append(slc, []int{s[0]-1, s[1]})
    //         re = append(re, []int{s[0]-1, s[1]})
    //         arr[s[0]-1][s[1]] = true
    //     }
    //     if s[1]<C-1 && !arr[s[0]][s[1]+1] {
    //         slc = append(slc, []int{s[0], s[1]+1})
    //         re = append(re, []int{s[0], s[1]+1})
    //         arr[s[0]][s[1]+1] = true
    //     }
    //     if s[1]>0 && !arr[s[0]][s[1]-1] {
    //         slc = append(slc, []int{s[0], s[1]-1})
    //         re = append(re, []int{s[0], s[1]-1})
    //         arr[s[0]][s[1]-1] = true
    //     }
    // }
    // return re
    r,c = r0,c0
    re := make(v, 0)
    for i:=0;i<R;i++ {
        for j:=0;j<C;j++ {
            re = append(re, []int{i,j})
        }
    }
    fmt.Println(re)
    sort.Sort(re)
    return re
}
func (s v)Len() int {
    return len(s)
}
func (s v)Less(i, j int) bool {
    return abs(s[i][0]-r) + abs(s[i][1]-c) < abs(s[j][0]-r) + abs(s[j][1]-c)
}
func (s v)Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}
func abs(a int) int {
    if a < 0 {
        return -1 * a
    }
    return a
}