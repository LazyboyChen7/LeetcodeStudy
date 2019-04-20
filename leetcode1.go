func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,v := range nums {
        n := target - v
        if idx,ok := m[n]; ok {
            return []int{i, idx}
        }
        m[v] = i
    }
    return nil
}