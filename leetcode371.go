func getSum(a int, b int) int {
    for b != 0 {
        t := a^b
        b = (a&b) << 1
        a = t
    }
    return a
}
