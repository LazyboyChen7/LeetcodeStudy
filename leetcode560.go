func subarraySum(nums []int, k int) int {
	var sum,re int
	m := make(map[int]int)
	m[0] = 1
	for _, v := range nums {
		sum += v
		re += m[sum - k]
		m[sum]++
	}
	return re
}
