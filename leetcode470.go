func rand10() int {
    var i int = 100
    for i>40 {
        i = 7*(rand7()-1) + rand7()
    }
    return i%10+1
}