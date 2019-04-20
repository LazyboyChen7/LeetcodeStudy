func maxScoreSightseeingPair(A []int) int {
    var re int = 0
    var maxi int = 0
    for i:=0;i<len(A);i++ {
        if maxi + A[i] - i > re {
            re = maxi + A[i] - i
        }
        if i + A[i] > maxi {
            maxi = i + A[i]
        }
    }
    return re
}