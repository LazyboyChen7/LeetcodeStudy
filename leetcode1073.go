func addNegabinary(arr1 []int, arr2 []int) []int {
    re := make([]int,0)
    var t int = max(len(arr1),len(arr2))
    var flg int = 0
    for i:=0;i<t||flg!=0;i++ {
        r := flg
        if i<len(arr1) {
            r += arr1[len(arr1)-1-i]
        }
        if i<len(arr2) {
            r += arr2[len(arr2)-1-i]
        }
        if r%2 == 0 {
            re = append(re, 0)
            flg = r/-2
        } else {
            re = append(re, 1)
            flg = (r-1)/-2
        }
    }
    reverse(re)
    var idx int
    for re[idx]==0&&idx<len(re)-1 {
        idx++
    }
    re = re[idx:]
    return re
}

func reverse(s []int) {
    for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1 {
        s[i],s[j] = s[j],s[i]
    }
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
