// 1005. Maximize Sum Of Array After K Negations

func largestSumAfterKNegations(A []int, K int) int {
    sort.Ints(A)
    sum := 0
    i := 0
    for ;K>0; {
        if A[i] <= 0 {
            A[i] = -A[i]
        } else {
            if K%2 == 0 {
                break;
            }
            if i<1 {
                A[i] = -A[i]
                break
            }
            if(A[i] < A[i-1]) {
                A[i] = -A[i]
            } else {
                A[i-1] = -A[i-1]
            }
            break;
        }
        K--
        i++
    }
    for _,a := range A {
        sum += a
    }
    return sum
}

// 执行用时 : 8 ms, 在Maximize Sum Of Array After K Negations的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.7 MB, 在Maximize Sum Of Array After K Negations的Go提交中击败了100.00% 的用户

