func canReach(arr []int, start int) bool {
    visited := map[int]struct{}{start: struct{}{}}
    q := []int{start}
    n := len(arr)
    for len(q) != 0 {
        idx := q[0]
        q = q[1:]
        if arr[idx] == 0 {return true} 
        numJumps := arr[idx]
        forward := idx + numJumps
        backward := idx-numJumps
        
        if forward <= n-1 {
            if arr[forward] == 0 {
                return true
            }
            if _, ok := visited[forward]; !ok {
                q = append(q, forward)
                visited[forward] = struct{}{}
            }
        }
        
        if backward >= 0 {
            if arr[backward] == 0 {
                return true
            }
            if _, ok := visited[backward]; !ok {
                q = append(q, backward)
                visited[backward] = struct{}{}
            }
        }
        
    }
    
    return false
    
}

func abs(n int) int {
    if n < 0 {return n*-1}
    return n
}