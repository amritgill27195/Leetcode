func jump(nums []int) int {
    visited := map[int]struct{}{
        0: struct{}{},
    }
    q := []int{0}
    jumps := 0
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq >= len(nums)-1 {return jumps}

            for i := nums[dq]; i >= 1; i-- {
                nextIdx := dq + i
                if nextIdx >= len(nums)-1 {return jumps+1}
                _, ok := visited[nextIdx]
                if !ok {
                    visited[nextIdx] = struct{}{}
                    q = append(q, nextIdx)
                }
            }
            qSize--
        }
        jumps++
    }
    return jumps
} 