func preimageSizeFZF(K int) int {
    i,sum := 1,1
    for K > sum {
        i *= 5
        sum += i
    }
    if K<sum {
        sum -= i
        i /= 5
    }
    var diff int = K-sum
    for diff != 0 {
        if sum == 1 && K == 5 {
            return 0
        }
        if diff/sum == 5 || diff%sum == 5 {
            return 0
        }
        diff %= sum
        sum -= i
        i /= 5
    }
    return 5
}