// func climbStairs(n int) int {
    
//     dp := make([]int, n+1)
//     dp[0] = 1
//     dp[1] = 1
    
//     for i := 2; i < len(dp); i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
    
//     return dp[len(dp)-1]
    
// }



// we only need two vars
func climbStairs(n int) int {
    
    dp := []int{1,1}
    
    for i := 2; i <= n; i++ {
        oneStepBack := dp[0]
        twoStepBack := dp[1]
        dp[0] = twoStepBack
        dp[1] = oneStepBack+twoStepBack
    }
    
    return dp[len(dp)-1]
    
}