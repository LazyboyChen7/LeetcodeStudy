// 1004. Max Consecutive Ones III

func longestOnes(A []int, K int) int {
    re, low, high := 0, 0, 0
    for high < len(A) {
        if A[high] == 1 {
            high++
        } else {
            if K != 0 {
                K--;
                high++;
            } else { 
                if A[low] == 0 {
                    K++;
                    low++;
                } else {
                    low++
                }
            }
        }
        if high-low > re {
            re = high-low
        }
    }
    return re;
}