func canVisitAllRooms(rooms [][]int) bool {
    m := map[int]bool{0:true}
    slc := make([]int,0)
    for _,v := range rooms[0] {
        slc = append(slc, v)
        m[v] = true
    }
    for len(slc) != 0 {
        n := slc[0]
        slc = slc[1:]
        for _,v := range rooms[n] {
            if !m[v] {
                slc = append(slc, v)
                m[v] = true
            }
        }
    }
    if len(m) == len(rooms) {
        return true
    }
    return false
}