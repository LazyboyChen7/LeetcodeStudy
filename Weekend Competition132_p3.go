func longestArithSeqLength(A []int) int {
	n, res := len(A), 1
	arr := make([]map[int]int, n)
	for i := range arr {
		arr[i] = make(map[int]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			var diff int = A[i] - A[j]
			if arr[i][diff] < arr[j][diff]+1 {
				arr[i][diff] = arr[j][diff] + 1
			}
			if res < arr[i][diff] {
				res = arr[i][diff]
			}
		}
	}
	return res + 1
}