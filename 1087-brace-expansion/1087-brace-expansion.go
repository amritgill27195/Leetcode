func expand(s string) []string {
    sList := [][]string{}
    i := 0
    for i < len(s) {
        grp := []string{}
        char := s[i]
        if char == ',' {continue}
        if char == '{' {
            i++
            for s[i] != '}' {
                if s[i] == ',' {i++; continue}
                grp = append(grp, string(s[i]))
                i++
            }
        } else {
            grp = append(grp, string(s[i]))
        }
        i++
        sort.Strings(grp)
        sList = append(sList, grp)
    }
    result := []string{}
    var backtrack func(outterIdx int, path string)
    backtrack = func(outterIdx int, path string) {
        
        // base
        if outterIdx == len(sList) {
            result = append(result, path)
            return
        }

        // logic
        for i := 0; i < len(sList[outterIdx]); i++ {
            path += sList[outterIdx][i]
            backtrack(outterIdx+1, path)
            path = path[:len(path)-1]
        }
        
    }
    backtrack(0, "")
    return result
}