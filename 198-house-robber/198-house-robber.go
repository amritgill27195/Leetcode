

// func rob(nums []int) int {
//     max := math.MinInt64
//     var dfs func(start int, profit int)
//     dfs = func(start int, profit int) {
//         // base
//         if profit > max {
//             max = profit
//         }
//         if start >= len(nums) {
//             return
//         }
        
//         // logic
//         dfs(start+2, profit+nums[start])
//         dfs(start+1, profit)
//     }
    
//     dfs(0,0)
//     return max
// }


// top - down using dp matrix
func rob(nums []int) int {
    dp := make([][]int, len(nums))
    for idx, _ := range dp {
        dp[idx] = make([]int, 2) // 0 | 1 case
    }
    dp[0][0] = 0
    dp[0][1] = nums[0]
    for i := 1; i < len(dp); i++ {
        for j := 0; j < len(dp[0]); j++ {
            if j == 0 {
                dp[i][j] = int(math.Max(float64(dp[i-1][0]),float64(dp[i-1][1]))) 
            } else {
                dp[i][j] = nums[i] + dp[i-1][0]
            }
        }
    }
    if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
        return dp[len(dp)-1][0]
    }
    return dp[len(dp)-1][1]
}