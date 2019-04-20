func findLength(A []int, B []int) int {
    re := 0
    var dp [][]int
    for i:=0 ;i<=len(A); i++ {
        p := make([]int, len(A)+1)
        dp = append(dp, p)
    }
    for i := 1; i <= len(A); i++ {
        for j := 1; j <= len(B); j++ {
            if A[i-1] == B[j-1] {
                dp[i][j] = dp[i-1][j-1]+1
                if dp[i][j] > re {
                    re = dp[i][j]
                }
            }
        }
    }
    return re
}