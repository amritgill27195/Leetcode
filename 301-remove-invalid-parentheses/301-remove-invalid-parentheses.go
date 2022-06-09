func removeInvalidParentheses(s string) []string {
    result := []string{}
    maxLen := 0
    visited := map[string]struct{}{}
    var backtrack func(path string)
    backtrack = func(path string) {
        // base
        
        if _, ok := visited[path]; ok {return}
        
        // logic
        visited[path] = struct{}{}
        if isValid(path) {
            if len(path) > maxLen {
                result = []string{path}
                maxLen = len(path)
            } else if len(path) == maxLen {
                result = append(result, path)
            }
        }
        for i := 0; i < len(path); i++ {
            if path[i] != '(' && path[i] != ')' {continue}
            child := path[:i] + path[i+1:]
            if len(child) < maxLen {continue}
            backtrack(child)
        }
    }
    backtrack(s)
    return result
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