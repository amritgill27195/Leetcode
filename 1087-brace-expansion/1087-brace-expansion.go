func expand(s string) []string {
    blocks := [][]string{}
    i := 0
    for i < len(s) {
        block := []string{}
        if s[i] == '{' {
            i++
            for s[i] != '}' {
                if s[i] == ',' {i++; continue}
                block = append(block, string(s[i]))
                i++
            }
        } else {
            block = append(block, string(s[i]))
        }
        i++
        sort.Strings(block)
        blocks = append(blocks, block)
        
    }
    result := []string{}
    var backtrack func(path string, start int)
    backtrack = func(path string, start int) {
        // base
        if start == len(blocks) {
            result = append(result, path)
            return
        }
        
        // logic
        for i := 0; i < len(blocks[start]); i++ {
            path += blocks[start][i]
            backtrack(path, start+1)
            path = path[:len(path)-1]
        }
    }
    backtrack("", 0)
    return result
}