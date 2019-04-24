type TrieNode struct {
    flg bool
    child map[byte]*TrieNode
}
func newTrieNode() *TrieNode {
    return &TrieNode{flg:false, child:make(map[byte]*TrieNode)}
}

type StreamChecker struct {
    root *TrieNode
    letter []byte
    maxLen int
}

func NewChecker() StreamChecker {
    return StreamChecker{root:newTrieNode(),letter:make([]byte,0),maxLen:0}
}

func reverse(c []byte) []byte {
    for i,j:=0,len(c)-1;i<j;{
        c[i],c[j] = c[j],c[i]
        i++
        j--
    }
    return c
}

func (sc *StreamChecker)insert(s string) {
    cs := []byte(s)
	cs = reverse(cs)
    node := sc.root
    for i:=0;i<len(s);i++ {
        _,ok := node.child[cs[i]]
        if !ok {
            node.child[cs[i]] = newTrieNode()
        }
        node = node.child[cs[i]]
    }
    node.flg = true
}

func Constructor(words []string) StreamChecker {
    sc := NewChecker()
    for _,v := range words {
        sc.insert(v)
        if len(v) > sc.maxLen {
            sc.maxLen = len(v)
        }
    }
    return sc
}


func (this *StreamChecker) Query(letter byte) bool {
    this.letter = append(this.letter, letter)
    node := this.root
    for i,ml:= len(this.letter)-1,1;i>=0 && ml<=this.maxLen;i-- {
		_,ok := node.child[this.letter[i]]
        if !ok {
            return false
		}
        node = node.child[this.letter[i]]
		if node.flg {
			return true
		}
	}
    return node.flg
}