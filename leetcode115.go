func numDistinct(s string, t string) int {
    arr := make([][]int, len(t)+1)
    for i, _ := range arr {
        arr[i] = make([]int, len(s)+1)
    }
    for i := 0; i < len(arr[0]); i++ {
        arr[0][i] = 1
    }
    for i := 0; i < len(t); i++ {
        for j := 0; j < len(s); j++ {
            if t[i] == s[j] {
                arr[i+1][j+1] = arr[i][j] + arr[i+1][j]
            } else {
                arr[i+1][j+1] = arr[i+1][j]
            }
        }
    }
    return arr[len(arr)-1][len(arr[0])-1]

//     下面这个是大佬写的
//     f := make([]int, len(t)+1)
//     f[0] = 1
//     for i := range s {
//         for j := len(t)-1; j>=0; j-- {
//             if s[i] == t[j] {
//                 f[j+1] += f[j]
//             }
//         }
//     }
//     return f[len(t)]
}