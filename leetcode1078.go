func findOcurrences(text string, first string, second string) []string {
    s := strings.Split(text, " ")
    re := make([]string,0)
    for i:=0;i<len(s)-2;i++ {
        if s[i] == first && s[i+1] == second {
            re = append(re, s[i+2])
        }
    }
    return re
}
