func removeInvalidParentheses(s string) []string {
    
    visited := map[string]bool{s: true}
    q := []string{s}
    result := []string{}
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if yes := isValid(dq); yes{
                result = append(result, dq)
            }
            
            // enqueue all the childs
            if len(result) == 0 {
                for i := 0; i < len(dq); i++ {
                    if dq[i] == '(' || dq[i] == ')' {
                        child := string(dq[0:i]) + string(dq[i+1:])
                        if _, ok := visited[child]; !ok {
                            visited[child] = true
                            q = append(q, child)
                        }
                    }
                }
            }
            qSize--
        }
        if len(result) != 0 {
            break
        }
    }
    return result
}

func isValid(s string) bool {
    count := 0
    for _, ele := range s {
        if ele == '(' {
            count++
        } else if ele == ')' {
            count--
        }
        if count < 0 {return false}
    }
    return count==0
}