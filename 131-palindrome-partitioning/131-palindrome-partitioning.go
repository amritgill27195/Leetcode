func partition(s string) [][]string {
    
    var result [][]string
    var dfs func(start int, paths []string)
    dfs = func(start int, paths []string){
        // base
        if start >= len(s) {
            newL := make([]string, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        
        // logic
        for i := start; i < len(s); i++ {
            subStr := string(s[start:i+1])
            if isPalindrome(subStr) {
                // action
                paths = append(paths, subStr)
                // recurse
                dfs(i+1, paths)
                // backtrack
                paths = paths[:len(paths)-1]
            }
        }   
    }
    
    dfs(0, nil)
    return result
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s) ; i++ {
        curChar := string(s[i])
        lastChar := string(s[len(s)-i-1])
        if curChar != lastChar {
            return false
        }
    }
    return true
}