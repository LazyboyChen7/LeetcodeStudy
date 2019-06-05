func makesquare(nums []int) bool {
    if len(nums) < 4 {
        return false
    }
    var re int
    for _,v := range nums {
        re += v
    }
    if re&3 != 0 {
        return false
    }
    sort.Ints(nums)
    return dfs(nums, 4, re>>2, 0, 0)
}

func dfs(s []int, dep,target,sum,idx int) bool {
    if dep == 1 {
        return true
    }
    if target == sum {
        return dfs(s, dep-1, target, 0, 0)
    }
    if sum > target {
        return false
    }
    for i:=idx;i<len(s);i++ {
        if s[i] < 0 {
            continue
        }
        s[i] = -s[i]
        if dfs(s, dep, target, sum-s[i],i+1) {
            return true
        }
        s[i] = -s[i]
    }
    return false
}
