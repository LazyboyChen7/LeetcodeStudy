//------    973. K Closest Points to Origin

func kClosest(points [][]int, K int) [][]int {
	if len(points) == K {
		return points
	}
	var v []int
	var re [][]int
	m := make(map[int]int)
	for i := 0; i < len(points); i++ {
		m[i] = points[i][0]*points[i][0] + points[i][1]*points[i][1]
		v = append(v, i)
	}
	findK(m, v, 0, len(v)-1, K)
	for i := 0; i < K; i++ {
		re = append(re, points[v[i]])
	}
	return re
}

func findK(m map[int]int, s []int, start, end, k int) {
	p := m[s[end]]
	l := start
	for i := start; i < end; i++ {
		if m[s[i]] < p {
			s[l], s[i] = s[i], s[l]
			l++
		}
	}
	s[l], s[end] = s[end], s[l]
	if k == l {
		return
	} else if l < k {
		findK(m, s, l+1, end, k)
	} else {
		findK(m, s, start, l-1, k)
	}
}

//执行用时 : 928 ms, 在K Closest Points to Origin的Go提交中击败了92.86% 的用户
//内存消耗 : 65.6 MB, 在K Closest Points to Origin的Go提交中击败了100.00% 的用户
