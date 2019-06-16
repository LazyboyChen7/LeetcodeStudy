func maxEqualRowsAfterFlips(matrix [][]int) int {
    m := make(map[string]int)
    a := make([]byte,len(matrix[0]))
    for i,_ := range a {
        a[i] = '0'
    }
    m[string(a)] = 0
    var re int = 1
    for _,v := range matrix {
        s := proc(v)
        var flg bool
        for i,_ := range m {
            if is(s,i) {
                m[i]++
                if m[i] > re {
                    re = m[i]
                }
                flg = true
            }
        }
        if !flg {
            m[s]++
        }
    }
    return re
}

func proc(a []int) string {
    var s string
    for _,v := range a {
        s += strconv.Itoa(v)
    }
    return s
}

func is(s1,s2 string) bool {
    if s1 == s2 {
        return true
    }
    if rev(s1) == s2 {
        return true
    }
    return false
}

func rev(s string) string {
    b := []byte(s)
    for i,_ := range b {
        if b[i] == 49 {
            b[i] = 48
        } else {
            b[i] = 49
        }
    }
    return string(b)
}
