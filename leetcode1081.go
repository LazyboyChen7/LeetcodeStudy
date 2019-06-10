func smallestSubsequence(text string) string {
    m := make(map[byte]int)
    f := make(map[byte]bool)
    s := make([]byte,0)
    for i:=0;i<len(text);i++ {
        m[text[i]]++
    }
    for i:=0;i<len(text);i++ {
        m[text[i]]--
        if f[text[i]] {
            continue
        }
        for len(s) > 0 && s[len(s)-1] > text[i] && m[s[len(s)-1]] > 0 {
            f[s[len(s)-1]] = false
            s = s[:len(s)-1]
        }
        s = append(s, text[i])
        f[text[i]] = true
    }
    return string(s)
}
