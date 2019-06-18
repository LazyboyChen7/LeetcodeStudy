func addStrings(num1 string, num2 string) string {
    l1,l2 := len(num1),len(num2)
    if l1 < l2 {
        l1,l2 = l2,l1
        num1,num2 = num2,num1
    }
    re := make([]byte, l1+1)
    i,j,flg,k := l1-1,l2-1,0,l1
    for i>-1 || j>-1 || flg>0 {
        n := flg
        if i>-1 {
            n += int(num1[i]-'0')
        }
        if j>-1 {
            n += int(num2[j]-'0')
        }
        re[k] = byte(n%10) + '0'
        flg = n/10
        i--
        j--
        k--
    }
    if re[0] == 0 {
        re = re[1:]
    }
    return string(re)
}