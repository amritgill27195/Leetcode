// BFS
// time: exponential
// space: o(n) for the queue
func jump(nums []int) int {
    visitedSet := map[int]bool{}
    
    q := []int{0}
    visitedSet[0] = true
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
            // here childs are the next indicies
            // each child is discovered by jumping from current idx + valAtThatIdx
            for i := 1; i <= nums[dq]; i++ {
                newIdx := dq+i
                if newIdx >= len(nums)-1 {return jumps+1}
                if _, ok := visitedSet[newIdx]; !ok {
                    visitedSet[newIdx] = true
                    q = append(q, newIdx)
                }
            }
            qSize--
        }
        jumps++
    }
    return jumps
}