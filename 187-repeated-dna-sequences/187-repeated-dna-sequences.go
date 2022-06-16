func findRepeatedDnaSequences(s string) []string {
    state := map[string]bool{}
    str := ""
    out := []string{}
    for right := 0; right < len(s); right++ {
        str += string(s[right])
        if len(str) == 10 {
            used, ok := state[str]
            if !ok {
                state[str] = false
            } else if !used {
                out = append(out, str)
                state[str] = true
            }
            str = str[1:]
        }
    }
    return out
}