// func change(amount int, coins []int) int {
//     var helper func(start int, target int) int
//     helper = func(start int, target int) int {
//         // base
//         if target == 0 {
//             return 1
//         }
        
//         if target < 0 || start == len(coins) {
//             return 0
//         }
        
        
//         // logic
//         choose := helper(start, target-coins[start])
//         notChoose := helper(start+1, target)
        
//         return choose+notChoose
//     }
    
//     return helper(0, amount)
// }


func change(amount int, coins []int) int {
    m := len(coins)
    n := amount
    
    dp := make([][]int, m+1)
    for i, _ := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i < m+1; i++ {
        dp[i][0] = 1
    }
    
    for i := 1; i < len(dp); i++ { // coins
        for j := 1; j < len(dp[0]); j++ { // amounts 
            coinVal := coins[i-1]
            am := j
            if coinVal > am {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-coinVal]
            }
        }
    }
    
    return dp[len(dp)-1][len(dp[0])-1]
}
