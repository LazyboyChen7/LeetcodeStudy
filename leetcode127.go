func ladderLength(beginWord string, endWord string, wordList []string) int {
    var flg bool
    for i:=0;i<len(wordList);i++ {
        if wordList[i] == endWord {
            flg = true
            break
        }
    }
    if !flg {
        return 0
    }
    var re int = 1
    var wordLen int = len(beginWord)
    slc := make([]string, 0)
    slc = append(slc, beginWord)
    for len(slc) != 0 {
        var size = len(slc)
        for size != 0 {
            s := slc[0]
            for i:=0;i<len(wordList);i++ {
                var count int = 0
                for j:=0;j<wordLen;j++ {
                    if s[j]-wordList[i][j] != 0 {
                        count++
                    }
                }
                if 1 == count {
                    if wordList[i] == endWord {
                        re++
                        return re
                    }
                    slc = append(slc, wordList[i])
                    wordList = append(wordList[:i],wordList[i+1:]...)
                    i--
                }
            }
            slc = slc[1:]
            size--
        }
        re++
    }
    return 0
}