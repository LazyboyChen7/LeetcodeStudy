type point struct {
    x,y int
}
type pts []point
func maxPoints(points [][]int) int {
    var re int
    slc := make(pts, 0)
    m := make(map[point]int)
    for _,v := range points {
        p := point{x:v[0], y:v[1]}
        if _, ok := m[p]; !ok {
            slc = append(slc, p)
        }
        m[p]++
    }
    if len(slc) < 3 {
        for _,v := range m {
            re += v
        }
        return re
    }
    re = 0
    for i:=0;i<len(slc)-2;i++ {
        for j:=i+1;j<len(slc)-1;j++ {
            tmp := m[slc[i]] + m[slc[j]]
            for k:=j+1;k<len(slc);k++ {
                if isLine(slc[i],slc[j],slc[k]) {
                    tmp += m[slc[k]]
                }
            }
            re = max(re, tmp)
        }
    }
    return re
}

func isLine(a,b,c point) bool {
    var k1,k2 int64
    k1 = int64(a.x-b.x) * int64(a.y-c.y)
    k2 = int64(a.x-c.x) * int64(a.y-b.y)
    if k1 == k2 {
        return true
    }
    return false
}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a
}
