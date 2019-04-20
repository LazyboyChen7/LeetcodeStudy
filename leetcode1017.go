func baseNeg2(N int) string {
    if N == 0{
        return "0"
    }
    a := make([]int,0)
    var bit int
    var re string
    for N!=0 {
        bit = abs(N%(-2))
        a = append(a, bit)
        N = (N-bit)/-2
    }
    for len(a) != 0 {
        re += strconv.Itoa(a[len(a)-1])
        a = a[:len(a)-1]
    }
    return re
}

func abs(a int) int {
    if a>=0 {
        return a
    }
    return 0-a
}