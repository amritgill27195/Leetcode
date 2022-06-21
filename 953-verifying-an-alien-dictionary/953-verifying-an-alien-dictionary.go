func isAlienSorted(words []string, order string) bool {
    orderIdxMap := map[byte]int{}
    for i := 0; i < len(order); i++ {
        orderIdxMap[order[i]] = i
    }
    
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        if !isSorted(word1, word2, orderIdxMap) {
            return false
        }
    }
    return true
}

func isSorted(word1, word2 string, orderIdxMap map[byte]int) bool {
    w1 := 0
    w2 := 0
    for w1 < len(word1) && w2 < len(word2) {
        w1Char := word1[w1]
        w2Char := word2[w2] 
        if w1Char != w2Char {
            return orderIdxMap[w1Char] < orderIdxMap[w2Char]
        }
        w1++
        w2++
    }
            if len(word1) > len(word2) {return false }
    return true
}