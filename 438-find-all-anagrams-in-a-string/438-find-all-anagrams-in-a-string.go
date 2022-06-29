func findAnagrams(s string, p string) []int {
    pMap := map[byte]int{}
    for i := 0; i < len(p); i++ {
        pMap[p[i]]++
    }
    out := []int{}
    count := len(pMap)
    left := 0
    for right := 0; right < len(s); right++ {
        _, ok := pMap[s[right]]
        if ok {
            pMap[s[right]]--
            if val := pMap[s[right]]; val == 0 {
                count--
            }
        }
        if right-left+1 == len(p) {
            if count == 0 {
                out = append(out, left)
            }
            leftChar := s[left]
            if val, exists := pMap[leftChar]; exists {
                pMap[leftChar]++
                if val == 0 {count++}
            }
            left++
        }
    }
    return out
}