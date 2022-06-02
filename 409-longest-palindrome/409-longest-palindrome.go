func longestPalindrome(s string) int {
    freqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freqMap[s[i]]++
    }
    
    maxLen := 0
    for char, count := range freqMap {
        if count % 2 == 0 {
            maxLen+=count
            delete(freqMap, char)
        } else {
            evenCount := count-1
            maxLen += evenCount
            freqMap[char] = count-evenCount
        }
    }
    
    if len(freqMap) != 0 {maxLen++}
    return maxLen
}