// time : o(n)
// space : o(n)
// func climbStairs(n int) int {
    
//     dp := make([]int, n+1)
//     dp[0] = 1
//     dp[1] = 1
    
//     for i := 2; i < len(dp); i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
    
//     return dp[len(dp)-1]
    
// }



// we only need previous values
// how would we store the new total and where?
// well it turns out that the new total is the new twoStepsBack value
// but first promote the twoStepBack value to 1 step back position
// and then store the new total in twoStepBack position
// This way our space is constant

// time : o(n)
// space : o(1)
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