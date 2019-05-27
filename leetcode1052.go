func maxSatisfied(customers []int, grumpy []int, X int) int {
    slc := make([]int, len(customers))
    copy(slc, customers)
    for i,_ := range slc {
        slc[i] *= grumpy[i]
    }
    sum := 0
    for i:=0;i<X;i++ {
        sum += slc[i]
    }
    re := sum
    for i:=X;i<len(slc);i++ {
        sum -= slc[i-X]
        sum += slc[i]
        if re < sum {
            re = sum
        }
    }
    for i:=0;i<len(customers);i++ {
        if grumpy[i] == 0 {
            re += customers[i]
        }
    }
    return re
}