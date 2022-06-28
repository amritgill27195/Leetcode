func findAnagrams(s string, p string) []int {
    left := 0
    wState := map[byte]int{}
    pMap := map[byte]int{}
    for i := 0; i < len(p); i++ {
        pMap[p[i]]++
    }
    out := []int{}
    for right := 0; right < len(s); right++ {
        wState[s[right]]++
        if right-left+1 == len(p) {
            wStateIsAnagram := true
            for k, v := range wState {
                pMapVal, exists := pMap[k]
                if (!exists && v != 0 )|| pMapVal != v {
                    wStateIsAnagram = false
                    break
                } 
            }
            if wStateIsAnagram {
                out = append(out, left)
            }
            leftChar := s[left]
            wState[leftChar]--
            // if wState[leftChar] == 0 {delete(wState, leftChar)}
            left++
        }
    }
    return out
}