const mod int64 = 1000000007
func nthMagicalNumber(N int, A int, B int) int {
    a,b,n := int64(A),int64(B),int64(N)
    if a > b {
        a,b = b,a
    }
    if b%a == 0 {
        return int(a*n%mod)
    }
    var l,r int64 = a, a*n
    var mb int64 = a/proc(a,b)*b
    if r < b {
        return int(r%mod)
    }
    for l <= r {
        m := (r-l)/2 + l
        if m/a + m/b - m/mb < n {
            l = m + 1
        } else if (m%a == 0 || m%b == 0) && (m/a + m/b - m/mb == n){
            return int(m%mod)
        } else {
            r = m - 1
        }
    }
    return 0
}

func proc(a,b int64) int64 {
    for a != b {
        if a>b {
            a -= b
        } else {
            b -= a
        }
    }
    return a 
}