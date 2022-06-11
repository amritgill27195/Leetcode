// func rob(nums []int) int {
//     var dfs func(start int, p int) int
//     dfs = func(start,p int) int {
//         // base
//         if start >= len(nums) {
//             return p
//         }
        
//         // logic
//         choose := dfs(start+2, p+nums[start])
//         notChoose := dfs(start+1, p)
        
//         if choose > notChoose {
//             return choose
//         }
//         return notChoose
//     }
    
//     return dfs(0,0)
// }
    
// func rob(nums []int) int {
//     profit := 0
//     var dfs func(start int, p int)
//     dfs = func(start,p int) {
//         // base
//         if start >= len(nums) {
//             if p > profit {
//                 profit = p
//             }
//             return
//         }
//         dfs(start+2, p+nums[start])
//         dfs(start+1, p)
//     }
    
//     dfs(0,0)
//     return profit
// }

// top down
func rob(nums []int) int {
    dp := make([][]int, len(nums))
    for idx, _ := range dp {
        dp[idx] = make([]int, 2)
        if idx == 0 {
            dp[0][0] = 0
            dp[0][1] = nums[0]
        }
    }
    
    for i := 1; i < len(nums); i++ {
        dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1])))
        dp[i][1] = nums[i] + dp[i-1][0]
    }
    if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
        return dp[len(dp)-1][0]
    }
    return dp[len(dp)-1][1]
}
    