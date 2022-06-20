func isAlienSorted(words []string, order string) bool {
    charIdxMap := map[byte]int{}
    for i := 0; i < len(order); i++ {
        charIdxMap[order[i]] = i
    }
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        if !isSorted(charIdxMap, word1, word2) {
            return false
        }
    }
    return true 
}

func isSorted(charIdxMap map[byte]int, word1, word2 string) bool {
    for i := 0; i < len(word1) && i < len(word2); i++ {
        word1Char := word1[i]
        word2Char := word2[i]
        if word1Char != word2Char {
            return charIdxMap[word1Char] < charIdxMap[word2Char]
        }
    }
    
    // if we got here, the characters are sorted
    // but the smaller len must be word1 for word1 to be placed before word2
    if len(word1) > len(word2) {
        return false
    }
    
    return true
}