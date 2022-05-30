func findRepeatedDnaSequences(s string) []string {
    windowState := map[string]bool{}
    tmp := ""
    result := []string{}
    for right := 0; right < len(s); right++ {
        rightChar := string(s[right])
        tmp += rightChar
        if len(tmp) == 10 {
            val , ok := windowState[tmp]
            if !ok {
                windowState[tmp] = false
            } else if val == false {
                windowState[tmp] = true
                result = append(result, tmp)
            }
            tmp = tmp[1:]
        }
    }
    return result
}