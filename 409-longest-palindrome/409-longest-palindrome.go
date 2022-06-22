func longestPalindrome(s string) int {
    maxLen := 0
    freqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freqMap[s[i]]++
    }
    
    for k, v := range freqMap {
        if v % 2 == 0 {
            maxLen += v
            delete(freqMap, k)
        } else if v != 1 && (v-1) % 2 == 0 {
            maxLen += v-1
            freqMap[k] = v - (v-1)
        }
    }
    
    if len(freqMap) != 0 {
        maxLen++
    }
        
    return maxLen
}