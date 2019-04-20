func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    var re string
    for i:=0;i<len(strs[0]);i++ {
        for j:=0;j<len(strs);j++ {
            if i>=len(strs[j]) || strs[j][i] != strs[0][i] {
                return re
            }
        }
        re += string(strs[0][i])
    }
    return re
}