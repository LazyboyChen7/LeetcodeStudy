func smallestRepunitDivByK(K int) int {
    m := make(map[int]bool)
    var re int = 1;
    var prod = 1%K;
    for !m[prod] {
        if(prod == 0) {
            return re;
        }
        m[prod] = true;
        prod = (prod*10+1)%K;
        fmt.Println(prod)
        re++;
    }
    return -1;
}