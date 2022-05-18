// func jump(nums []int) int {
//     visited := map[int]struct{}{
//         0: struct{}{},
//     }
//     q := []int{0}
//     jumps := 0
    
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if dq >= len(nums)-1 {return jumps}

//             for i := nums[dq]; i >= 1; i-- {
//                 nextIdx := dq + i
//                 if nextIdx >= len(nums)-1 {return jumps+1}
//                 _, ok := visited[nextIdx]
//                 if !ok {
//                     visited[nextIdx] = struct{}{}
//                     q = append(q, nextIdx)
//                 }
//             }
//             qSize--
//         }
//         jumps++
//     }
//     return jumps
// } 

/*
    approach: greedy
    - Using 2 intervals
    - Current interval = nums[0]
    - Next interval = num[0]
    - We will only increase our jump count once we have reached the current interval
    - While our i is headed towards current interval,
        - we will look for the next number that makes us jump the farthest distance possible
        - and we will maintain that with nextInterval
    - Once our i has reached the current interval idx, then we will increase our jump count by 1
    - and we will promote our next interval to current interval -- i.e do not increase the jump count until i == current interval and i has not reached the end
    - Finally return the jump count
    
    time: o(n)
    space: 0(1)

*/
func jump(nums []int) int {
    currInt := nums[0]
    nextInt := nums[0]
    jumps := 1
    if len(nums) < 2 {return 0}

    for i := 1; i < len(nums); i++ {
        nextInt = int(math.Max(float64(nextInt), float64(i+nums[i])))
        if i == currInt && i != len(nums)-1 {
            jumps++
            currInt = nextInt
        }
    }
    return jumps
}