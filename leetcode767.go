func reorganizeString(S string) string {
    var mx int
    s := []byte(S)
    m := make(map[byte]int)
    for _,v := range s{
        m[v]++
        if m[v] > mx {
            mx = m[v]
        }
    }
    if mx > (len(s)+1) / 2 {
        return ""
    }
    slc := make([]pt, 0)
    for k, v := range m {
        slc = append(slc, pt{k, v})
    }
    sort.Slice(slc, func(i, j int) bool {
        return slc[i].n > slc[j].n
    })
    re := make([]byte, len(s))
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
    return string(re)
}

type pt struct {
    v byte
    n int
}
