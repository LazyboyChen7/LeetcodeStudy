func numBusesToDestination(routes [][]int, S int, T int) int {
    if S == T {
        return 0
    }
    hadsta := make(map[int]bool)
    hadbus := make(map[int]bool)
    sta := make(map[int][]int)
    bus := make(map[int][]int)
    list := make([]int,0)
    for i:=0;i<len(routes);i++ {
        bus[i] = make([]int,0)
        for j:=0;j<len(routes[i]);j++ {
            if _,ok := sta[routes[i][j]]; !ok {
                sta[routes[i][j]] = make([]int,0)
            }
            if has(bus[i], routes[i][j]) {
                continue
            }
            sta[routes[i][j]] = append(sta[routes[i][j]], i)
            if routes[i][j] == S {
                hadsta[j] = true
                list = append(list, i)
            }
            bus[i] = append(bus[i], routes[i][j])
        }
    }
    var re int = 0
    for len(list) != 0{
        size := len(list)
        re++
        for size > 0{
            size--
            a := list[0]
            list = list[1:]
            if hadbus[a] {
                continue
            }
            hadbus[a] = true
            for _,v := range bus[a] {
                if hadsta[v] {
                    continue
                }
                if v == T {
                    return re
                }
                hadsta[v] = true
                for _,vl := range sta[v] {
                    list = append(list, vl)
                }
            }
        }
    }
    return -1
}

func has(a []int, b int) bool {
    for _,v := range a {
        if v == b {
            return true
        }
    }
    return false
}