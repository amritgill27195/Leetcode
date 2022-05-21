func partition(s string) [][]string {
    result := [][]string{}
    var dfs func(paths []string, startIdx int) 
    dfs = func(paths []string, startIdx int) {
        // base
        if startIdx == len(s) {
            newL := make([]string, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        
        
        // logic
        for i := startIdx; i < len(s); i++ {
            // action
            subStr := string(s[startIdx:i+1])
            if isValid(subStr) {
                paths = append(paths, subStr)
                // recurse
                dfs(paths, i+1)                
                // backtrack
                paths = paths[:len(paths)-1]
            }
            
        }
    }
    dfs(nil, 0)
    return result
}


func isValid(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        first := string(s[i])
        last := string(s[len(s)-1-i])
        if first != last {
            return false
        }
    }
    return true
}