// func removeInvalidParentheses(s string) []string {
    
//     result := []string{}
//     visited := map[string]struct{}{}
//     q := []string{s}
//     visited[s] = struct{}{}
    
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if isValid(dq) {
//                 result = append(result, dq)
//             }
//             for i := 0; i < len(dq); i++ {
//                 if dq[i] == '(' || dq[i] == ')' {
//                     child := string(dq[0:i]) + string(dq[i+1:])
//                     if _, ok := visited[child]; !ok {
//                         q = append(q, child)
//                         visited[child] = struct{}{}
//                     }
//                 }
//             }
//             qSize--
//         }
//         if len(result) > 0 {
//             break
//         }
//     }
//     return result
// }



func removeInvalidParentheses(str string) []string {
    result := []string{}
    max := 0
    visited := map[string]struct{}{}
    var dfs func(s string)
    dfs = func(s string) {
        // base
        if len(s) < max {return}
        if _, ok := visited[s]; ok {return}
        if isValid(s) {
            if len(s) > max {
                result = []string{s}
                max = len(s)
            } else if len(s) == max {
                result = append(result, s)
            }
            visited[s] = struct{}{}
            return
        }
        
        // logic
        visited[s] = struct{}{}
        for i := 0; i < len(s); i++ {
            if s[i] == '(' || s[i] == ')' {
                child := string(s[0:i]) + string(s[i+1:])
                dfs(child)
            }
        }
    }
    dfs(str)
    return result
}

func isValid( s string) bool {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
        }
        if count < 0 {
            return false
        }
    }
    return count == 0
}