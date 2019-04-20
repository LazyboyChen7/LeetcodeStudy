func prefixesDivBy5(A []int) []bool {
    var re int
    res := make([]bool,0)
    for i:=0;i<len(A);i++ {
        re = re << 1 + A[i]
        res = append(res, re%5==0)
        re %= 5
    }
    return res
}