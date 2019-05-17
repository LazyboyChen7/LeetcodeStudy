type MyCalendar struct {
    arr [][2]int
}


func Constructor() MyCalendar {
    arr := make([][2]int,0)
    return MyCalendar{arr : arr}
}


func (this *MyCalendar) Book(start int, end int) bool {
    slc := [2]int{start,end}
    for i,_ := range this.arr {
        if start >= this.arr[i][1] || end <= this.arr[i][0] {
            continue
        } else {
            return false
        }
    }
    this.arr = append(this.arr, slc)
    return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */