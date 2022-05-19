func partitionLabels(s string) []int {
    charLastIdx := map[rune]int{}
    for idx, char := range s {
        charLastIdx[char] = idx
    }
    result := []int{}
    
    start := -1
    windowEnd := -1
    for idx, char := range s {
        if start == -1 {
            start = idx
            windowEnd = charLastIdx[char]
        }
        if idx < windowEnd {
            lastIdxOfThisChar := charLastIdx[char]
            if lastIdxOfThisChar > windowEnd {
                windowEnd = lastIdxOfThisChar
            }
        } else {
            if idx == windowEnd {
                result = append(result, windowEnd-start+1)
                start = -1
                windowEnd = -1
            }
        }
    }
    return result
    
    return nil
}