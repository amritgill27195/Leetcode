func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1 // no way a bigger word can exist in a smaller word
    }
    if len(needle) == 0 {return 0}
    for i := 0; i < len(haystack)+1-len(needle); i++ {
        for j := 0; j < len(needle); j++ {
            if haystack[i+j] != needle[j] {
                break
            }
            if j == len(needle)-1 {
                return i
            }
        }
    }
    return -1
}