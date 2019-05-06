var re int
func findTargetSumWays(nums []int, S int) int {
    re = 0
    proc(0, 0, S, nums)
    return re
}

func proc(n, i, res int, slc []int) {
    if i == len(slc) {
        if n == res {
            re++
        }
        return
    }    
    proc(n+slc[i], i+1, res, slc)
    proc(n-slc[i], i+1, res, slc)
}