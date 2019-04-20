// 1007. Minimum Domino Rotations For Equal Row

func minDominoRotations(A []int, B []int) int {
    m := make(map[int]int)
    for _,a := range A {
        m[a]++
    }
    for _,b := range B {
        m[b]++
    }
    maxv, index := 0, 0
    for k,v := range m {
        if v > maxv {
            maxv = v
            index = k
        }
    }
    if maxv<len(A) {
        return -1
    }
    ca, cb := 0, 0
    for i:= 0;i<len(A);i++ {
        if A[i] != index && B[i] != index {
            return -1
        }
        if A[i] == index && B[i] == index {
            continue
        }
        if A[i] != index {
            ca++
        } else {
            cb++
        }
    }
    if ca > cb {
        return cb
    }
    return ca
}

