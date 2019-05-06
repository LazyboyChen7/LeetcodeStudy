func PredictTheWinner(nums []int) bool {
    dp := make([][]int, len(nums))
    for i,_ :=range dp {
        dp[i] = make([]int, len(nums))
    }
    dp[len(nums)-1][len(nums)-1] = nums[len(nums)-1]
    for l:=len(nums)-2;l>=0;l-- {
        for r := l;r<len(dp);r++ {
            if l == r {
                dp[l][r] = nums[l]
            } else {
                dp[l][r] = max(nums[l]-dp[l+1][r], nums[r]-dp[l][r-1])
            }
        }
    }
    return dp[0][len(nums)-1] >= 0
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}