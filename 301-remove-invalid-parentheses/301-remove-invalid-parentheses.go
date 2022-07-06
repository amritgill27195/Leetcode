func removeInvalidParentheses(s string) []string {    
    maxLen := 0
    out := []string{}
    visited := map[string]struct{}{}
    
    var dfs func(path string)
    dfs = func( path string) {
        // base
        if _, ok := visited[path]; ok { return }
        visited[path] = struct{}{}

        // logic
        if isValid(path) {
            if len(path) > maxLen {
                out = []string{path}
                maxLen = len(path)
            } else {
                out = append(out, path)
            }
            return
        }
        
        for i := 0 ; i < len(path); i++ {
            // if path[i] == '(' || path[i] == ')' { 
                child := string(path[0:i]) + string(path[i+1:])
                if len(child) < maxLen{continue}
                dfs(child)
            // }
        }
    }
    
    dfs(s)
    if len(out) == 0 {
        out = []string{""}
    }
    return out
}

func isValid(s string) bool {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {count++}
        if s[i] == ')' {count--}
        if count < 0 {
            return false
        }
    }
    return count == 0
}