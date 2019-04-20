// 485. Max Consecutive Ones

func findMaxConsecutiveOnes(nums []int) int {
    max, temp := 0, 0
	for _, v := range nums {
		if v == 1 {
			temp++
		} else {
				temp = 0
		}

		if temp > max {
			max = temp
		}
	}
	return max
}