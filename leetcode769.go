func maxChunksToSorted(arr []int) int {
    var re int
    num := arr[0]
    for i:=0;i<len(arr);i++ {
        if arr[i]>num {
            num = arr[i]
        }
        if num == i {
            re++
            if i == len(arr)-1 {
                return re
            } else {
                num = arr[i+1]
            }
        }
    }
    return re
}