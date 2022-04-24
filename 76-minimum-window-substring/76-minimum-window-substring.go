func minWindow(s string, t string) string {
    tMap := map[string]int{}
    for _, char := range t {
        tMap[string(char)]++
    }
    count := 0
    left := 0
    startIdx := 0
    endIdx := len(s)+1
    for right := 0; right < len(s); right++ {
        rightChar := string(s[right])
        if _, ok := tMap[rightChar]; ok {
            tMap[rightChar]--
            if val := tMap[rightChar]; val == 0 {
                count++
            }
        }
        for count == len(tMap) {
            windowSize := right-left+1
            if windowSize < endIdx-startIdx+1 {
                startIdx = left
                endIdx = right
            }
            leftChar := string(s[left])
            if val, ok := tMap[leftChar]; ok {
                tMap[leftChar]++
                if val == 0 {
                    count--
                }
            }
            left++
        }
    }
    if endIdx == len(s)+1 {return "" }
    return string(s[startIdx:endIdx+1])
}