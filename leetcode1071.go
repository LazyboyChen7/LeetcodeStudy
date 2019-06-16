func gcdOfStrings(str1 string, str2 string) string {
    if len(str1) == len(str2) {
        if str1 == str2 {
            return str1
        } else {
            return ""
        }
    }
    s1,s2 := str1,str2
    for len(s1) != len(s2) {
        if len(s1) < len(s2) {
            s1,s2 = s2,s1
        }
        if !isSame(s1,s2) {
            return ""
        } else {
            s1 = s1[len(s2):]
        }
    }
    if s1 == s2 {
        return s1
    }
    return ""
}

func isSame(a,b string) bool {
    if a[:len(b)] == b {
        return true
    }
    return false
}
