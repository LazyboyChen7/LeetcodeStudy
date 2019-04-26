func validMountainArray(A []int) bool {
    if len(A) < 3 || A[0] > A[1] {
		return false
	}
	down := false
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		} else if A[i] < A[i-1] {
			down = true
		} else {
			if down {
				return false
			}
		}
	}
	return down
}