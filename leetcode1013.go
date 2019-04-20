func canThreePartsEqualSum(A []int) bool {
    var sum int
    for _,v := range A {
        sum += v
    }
    if sum%3 != 0 {
        return false
    }
    n1,n2,i,j := 0,0,0,len(A)-1
    for i<j {
        if n1 != sum/3 {
            n1 += A[i]
            i++
        }
        if n2 != sum/3 {
            n2 += A[j]
            j--
        }
        if n1 == sum/3 && n2 == sum/3 {
            return true
        }
    }
    return false
}