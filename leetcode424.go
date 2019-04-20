leetcode 424

func characterReplacement(s string, k int) int {
    var res, maxCnt, start int
    slc := make([]int, 26)
    for i:=0;i<len(s);i++ {
        slc[s[i] - 'A']++
        if slc[s[i] - 'A'] > maxCnt {
            maxCnt = slc[s[i] - 'A']
        }
        for i - start + 1 - maxCnt > k {
            slc[s[start] - 'A']--
            start++
        }
        if i - start + 1 > res {
             res = i - start + 1
        }
    }
    return res
}