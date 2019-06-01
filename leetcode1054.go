func rearrangeBarcodes(barcodes []int) []int {
    m := make(map[int]int)
    for _,v := range barcodes{
        m[v]++
    }
    slc := make([]pt, 0)
    for k, v := range m {
        slc = append(slc, pt{k, v})
    }
    sort.Slice(slc, func(i, j int) bool {
        return slc[i].n > slc[j].n
    })
    re := make([]int, len(barcodes))
    start := 0
    for _, s := range slc {
        v,n := s.v, s.n
        for n > 0 {
            if re[start] == 0 {
                re[start] = v
                start += 2
                n--
            } else {
                start++
            }
            start = start%len(re)
        }
    }
    return re
}

type pt struct {
    v int
    n int
}
