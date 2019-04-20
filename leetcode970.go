func powerfulIntegers(x int, y int, bound int) []int {
    var res []int
    m := make(map[int]bool)
    for a := 1; a <= bound; a *= x {
        for b := 1; a + b <= bound; b *= y {
            num := a + b
            if !m[num] {
                m[num] = true;
                res = append(res, num)
            }
            if y == 1 {
                break;
            }
        }
        if x == 1 {
            break;
        }
    }
    return res;
}
