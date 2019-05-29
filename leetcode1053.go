func prevPermOpt1(A []int) []int {
    if len(A) == 1 {
        return A
    }
    idx1 := -1
    for i:=len(A)-1;i>0;i-- {
        if A[i] < A[i-1] {
            idx1 = i-1
            break
        }
    }
    if idx1 == -1 {
        return A
    }
    idx2 := idx1 + 1
    sum := -1
    for i:=idx2;i<len(A);i++ {
        if A[i]<A[idx1] && A[i] > sum {
            idx2 = i 
            sum = A[i]
        }
    }
    A[idx1],A[idx2] = A[idx2],A[idx1]
    return A
}