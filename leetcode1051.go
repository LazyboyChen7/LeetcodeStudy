func heightChecker(heights []int) int {
    slc := make([]int, len(heights))
    copy(slc, heights)
    sort.Ints(slc)
    var re int
    for i,_ := range slc {
        if slc[i] != heights[i] {
            re++
        }
    }
    return re
}