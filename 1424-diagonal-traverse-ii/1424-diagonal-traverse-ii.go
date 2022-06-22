func findDiagonalOrder(nums [][]int) []int {
    diagGroup := map[int][]int{}
    m := len(nums)
    
    max := 0
    
    // time: o(mn)
    for i := m-1; i >= 0; i-- {
        for j := 0; j < len(nums[i]); j++ {
            rowColSum := i+j
            diagGroup[rowColSum] = append(diagGroup[rowColSum], nums[i][j])
            if rowColSum > max {max = rowColSum}
        }
    }
    out := []int{}
    
    // time: o(mn)
    for i := 0; i <= max; i++ {
        out = append(out, diagGroup[i]...)
    }
    
    // time: 2o(mn) = o(mn)
    // space: 
    
    return out
    
    
}