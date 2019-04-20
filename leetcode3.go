func lengthOfLongestSubstring(s string) int {
    var res, left, i int
    n := len(s)
    m := make(map[byte]int)
    for i = 0; i < n; i++ {
        if m[s[i]] > left {
            left = m[s[i]]
        }
        m[s[i]] = i + 1;
        if i-left+1 > res {
            res = i-left+1
        }
    }
    return res;
}