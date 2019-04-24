type cw struct {
    sum,l,r int
}
func maxSumTwoNoOverlap(A []int, L int, M int) int {
    cl := make([]cw, len(A)-L+1)
    cm := make([]cw, len(A)-M+1)
    var sm int 
    for i:=0;i<L;i++ {
        sm += A[i]
    }
    cl[0] = cw{sm,0,L-1}
    for i:=1;i<len(cl);i++ {
        cl[i] = cw{cl[i-1].sum-A[cl[i-1].l]+A[i+L-1],i,i+L-1}
    }
    sm = 0
    for i:=0;i<M;i++ {
        sm += A[i]
    }
    fmt.Println(cl)
    
    cm[0] = cw{sm,0,M-1}
    for i:=1;i<len(cm);i++ {
        cm[i] = cw{cm[i-1].sum-A[cm[i-1].l]+A[i+M-1],i,i+M-1}
    }
    fmt.Println(cm)
    
    var re int
    for i:=0;i<len(cl);i++ {
        for j:=0;j<len(cm);j++ {
            var sum int
            if cl[i].r < cm[j].l || cl[i].l>cm[j].r {
                sum = cl[i].sum + cm[j].sum
                if sum > re {
                    re = sum
                    fmt.Println(cl[i], cm[j])
                }
            }
        }
    }
    
    return re
}