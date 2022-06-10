

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
// func rob(nums []int) int {
//     dp := make([][]int, len(nums))
//     for idx, _ := range dp {
//         dp[idx] = make([]int, 2) // 0 | 1 case
//     }
//     dp[0][0] = 0
//     dp[0][1] = nums[0]
//     for i := 1; i < len(dp); i++ {
//         dp[i][0] = int(math.Max(float64(dp[i-1][0]),float64(dp[i-1][1]))) 
//         dp[i][1] = nums[i] + dp[i-1][0]
//     }
//     if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
//         return dp[len(dp)-1][0]
//     }
//     return dp[len(dp)-1][1]
// }


// bottom up using dp matrix
// func rob(nums []int) int {
//     dp := make([][]int, len(nums))
//     for idx, _ := range dp {
//         dp[idx] = make([]int, 2) // 0 | 1 case
//     }
//     dp[len(dp)-1][0] = 0
//     dp[len(dp)-1][1] = nums[len(nums)-1]
//     for i := len(dp)-2; i >= 0; i-- {
//         dp[i][0] = int(math.Max(float64(dp[i+1][0]),float64(dp[i+1][1]))) 
//         dp[i][1] = nums[i] + dp[i+1][0]
//     }
//     return int(math.Max(float64(dp[0][1]),float64(dp[0][0])))
// }


func rob(nums []int) int {
    dp := make([]int, 2)
    dp[0] = 0
    dp[1] = nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        choose := dp[1]
        notChoose := dp[0]
        
        dp[0] = int(math.Max(float64(choose),float64(notChoose))) 
        dp[1] = nums[i] + notChoose
    }
    return int(math.Max(float64(dp[0]),float64(dp[1])))
}