var slc []int
var re int
func numTilePossibilities(tiles string) int {
    slc = make([]int, 26)
    re = 0
    for i:=0;i<len(tiles);i++ {
        slc[tiles[i]-'A']++
    }
    proc()
    return re
}

func proc() {
    for i:=0;i<26;i++ {
        if slc[i] > 0 {
            re++
            slc[i]--
            proc()
            slc[i]++
        }
    }
}
