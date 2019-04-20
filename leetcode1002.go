// 1002. Find Common Characters

func commonChars(A []string) []string {
    var re []string
    if len(A) < 1 {
        return re
    }
    m := make(map[rune]int)
    for _,c := range A[0] {
        m[c]++
    }
    for k,v := range m {
        a := string(k)
        for j := 1; j < len(A); j++ {
	   n := strings.Count(A[j], a)
	   if n < v {
	       v = n
	   }
        }
	for j := 0; j < v; j++ {
	   re = append(re, a)
	}
    }
    return re
}

// 执行用时 : 4 ms, 在Find Common Characters的Go提交中击败了100.00% 的用户
// 内存消耗 : 3.2 MB, 在Find Common Characters的Go提交中击败了100.00% 的用户