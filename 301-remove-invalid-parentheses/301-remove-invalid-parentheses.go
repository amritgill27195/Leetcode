func removeInvalidParentheses(s string) []string {    
    out := []string{}
    q := []string{s}
    visited := map[string]bool{s: true}    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if isValid(dq) {
                out = append(out, dq)
            } else {
                for i := 0; i < len(dq); i++ {
                    child := string(dq[0:i] + dq[i+1:])
                    if _, ok := visited[child]; !ok {
                        visited[child] = true
                        q = append(q, child)
                    }
                }
            }
            qSize--
        }
        if len(out) != 0 {
            break
        }
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