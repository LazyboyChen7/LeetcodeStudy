func queryString(S string, N int) bool {
    for i:=1;i<=N;i++ {
        var num int
        var s string
        for j:=i;j>0;j=j/2 {
            if j%2 == 0 {
                num = 0
            } else {
                num = 1
            }
            s = strconv.Itoa(num) + s
        }
        if !strings.Contains(S, s) {
            return false
        }
    }
    return true
}