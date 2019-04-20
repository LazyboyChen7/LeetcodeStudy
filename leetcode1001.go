// type pt struct {
//     x,y int
// }
// 因为要把坐标存到map里，最开始使用struct，后来换成数组，但是系统判断并没有实质性的空间节约。

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
    r,c,x,y := make(map[int]int),make(map[int]int),make(map[int]int),make(map[int]int)
    m := make(map[[2]int]bool)
    for _,k := range lamps {
        qx,qy := k[0],k[1]
        r[qx] += 1
        c[qy] += 1
        x[qx+qy] += 1
        y[qx-qy] += 1
        var a = [2]int{qx, qy}
        m[a] = true
    }
    re := make([]int, len(queries))
    for i,k := range queries {
        qx,qy := k[0],k[1]
        if r[qx]!=0 || c[qy]!=0 || x[qx+qy]!=0 || y[qx-qy]!=0 {
            re[i] = 1
        } 
        for ix:=qx-1;ix<qx+2;ix++ {
            for iy:=qy-1;iy<qy+2;iy++ {
                var p = [2]int{ix,iy}
                if m[p] {
                    r[ix] -= 1
                    c[iy] -= 1
                    x[ix+iy] -= 1
                    y[ix-iy] -= 1
                    m[p] = false
                }
            }
        }
    }
    return re
}


// 下面这种方法 倒数第二个超时了
/*
func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
    m := make(map[int]int)
    for i:=0;i<len(lamps);i++ {
        m[i] = 1
    }
    re := make([]int,0)
    for i:=0;i<len(queries);i++ {
        re = append(re, cls(lamps, m, queries[i]))
    }
    return re
}

func cls(l [][]int, m map[int]int, q []int) int {
    var re int
    for i,k := range l {
        if m[i] == 1 && re == 0 && ((k[0] == q[0]) || (k[1] == q[1]) || 
            (k[0]+k[1] == q[0]+q[1]) || (q[0]-k[0]==q[1]-k[1])) {
            re = 1
            break
        }
    }
    for i,k := range l {
        if (q[0]-k[0])*(q[0]-k[0]) + (q[1]-k[1])*(q[1]-k[1]) <=2 {
            m[i] = 0
        }
    }
    
    return re
}*/
