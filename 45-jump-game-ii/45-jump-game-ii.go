func jump(nums []int) int {
    q := []int{0}
    visited := map[int]struct{}{}
    jumps := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {            
            dq := q[0]
            q = q[1:]
            if dq >= len(nums)-1 {
                return jumps
            }
            // explore all the childs
            for i := nums[dq]; i >= 1; i-- {
                nextIdx := i+dq
                if _, ok := visited[nextIdx]; ok {continue}
                visited[nextIdx] = struct{}{}
                if nextIdx >= len(nums)-1 {
                                return jumps+1
                            }
                q = append(q, nextIdx)
            }
            qSize--
        }
        jumps++
    }
    return jumps
}