func trap(height []int) int {
    if len(height)<2 {
        return 0
    }
    var maxn int = -1
    var maxIdx int = 0
    for i:=0;i<len(height);i++ {
        if height[i]>maxn {
            maxn = height[i]
            maxIdx = i
        }
    }
    var re int = 0
    var root int = height[0]
    for i:=0;i<maxIdx;i++ {
        if root<height[i]{
            root = height[i]
        } else {
            re += root - height[i]
        }
    }
    for i,root :=len(height)-1,height[len(height)-1];i>maxIdx;i-- {
        if root<height[i]{
            root = height[i]
        } else {
            re += root - height[i]
        }
    }
    return re
}