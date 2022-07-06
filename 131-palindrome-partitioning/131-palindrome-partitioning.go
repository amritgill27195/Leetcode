func partition(s string) [][]string {
    out := [][]string{}
    var dfs func(start int, paths []string)
    dfs = func(start int, paths []string) {
        
        // base
        if start == len(s) {
            newL := make([]string, len(paths))
            copy(newL, paths)
            out = append(out, newL)
            return
        }
        
        
        // logic
        for i := start; i < len(s); i++ {
            child := string(s[start:i+1])
            if isPalindrome(child) {
                // action
                paths = append(paths, child)
                // recurse
                dfs(i+1, paths)
                // backtrack
                paths = paths[:len(paths)-1]
            }
        }
    }
    dfs(0, nil)
    return out
}

func isPalindrome(s string) bool {
    left := 0
    right := len(s)-1
    for left < right {
        if s[left] == s[right] {
            left++
            right--
        } else{
            return false
        }
    }
    return true
}