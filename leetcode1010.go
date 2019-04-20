func numPairsDivisibleBy60(time []int) int {
    m := make(map[int]int)
    var re int
    for idx:=0;idx<len(time);idx++ {
        for i,v := range m {
            if v!=0 && (i+time[idx])%60 == 0 {
                re += v
                break
            } 
        }
        m[time[idx]%60]++
    }
    return re
}

// 这道题也可以遍历time，将time[i]%60存到map中，然后通过查找和为60的i与j，用map[i]与map[j]相乘 再累和得到结果。