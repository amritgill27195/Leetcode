func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sFreqMap := map[string]int{}
    
    for i := 0; i < len(s); i++ {
        char := string(s[i])
        sFreqMap[char]++
    }
    
    for i := 0; i < len(t); i++ {
        char := string(t[i])
        numTimes, exists := sFreqMap[char]
        if exists {
            sFreqMap[char]--
            if numTimes == 1 {
                delete(sFreqMap, char)
            }
        } else {
            return false
        }
    }
    
    return true
    
}