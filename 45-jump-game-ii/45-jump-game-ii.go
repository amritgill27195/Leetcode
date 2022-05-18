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