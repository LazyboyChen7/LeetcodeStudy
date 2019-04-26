func numsSameConsecDiff(N int, K int) []int {
    slc := []int{0,1,2,3,4,5,6,7,8,9}
    if N == 1 {
        return slc
    }
    slc = slc[1:]
    for N > 1 {
        N--
        arr := make([]int,0)
        for _,v := range slc {
            num := v%10
            if num-K>=0 && num-K<=9 {
                arr = append(arr, v*10+num-K)
            }
            if num+K!=num-K && num+K>=0 && num+K<=9 {
                arr = append(arr, v*10+num+K)
            }
        }
        slc = arr
        fmt.Println(slc)
    }
    return slc
}