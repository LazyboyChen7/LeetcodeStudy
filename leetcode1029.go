type point struct {
    idx, diff int
}
type points []point
func twoCitySchedCost(costs [][]int) int {
    var re int
    ps := make(points, len(costs))
    for i:=0;i<len(costs);i++ {
        pt := point{i,costs[i][0]-costs[i][1]}
        ps[i] = pt
    }
    sort.Sort(ps)
    for i,j:=0,len(ps)-1;i<j; {
        re += costs[ps[i].idx][0]
        re += costs[ps[j].idx][1]
        i++
        j--
    }
    
    return re
}

func (p points)Len() int {
    return len(p)
}
func (p points)Less(i,j int) bool {
    return p[i].diff < p[j].diff
}
func (p points)Swap(i,j int) {
    p[i],p[j] = p[j],p[i]
}