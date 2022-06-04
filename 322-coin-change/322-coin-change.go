// func coinChange(coins []int, amount int) int {
//     min := math.MaxInt64    
//     var dfs func(start int, count int, amount int)
//     dfs = func(start int, count int, amount int) {
//         // base
//         if amount == 0 {
//             if count < min {
//                 min = count
//             }
//             return
//         }
//         if amount < 0 || start == len(coins) {
//             return
//         }
        
        
//         // logic
//         dfs(start, count+1, amount-coins[start])
//         dfs(start+1, count, amount)
        
//     }
//     dfs(0, 0, amount)
//     if min == math.MaxInt64 {
//         return -1
//     }
//     return min
// }

// func coinChange(coins []int, amount int) int {
//     m := len(coins)
//     n := amount
    
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
    
//     for j := 1; j < len(dp[0]); j++ {
//         dp[0][j] = 10001 // max 
//     }
    
//     for i := 1; i < len(dp); i++ { // coin value
//         for j := 1; j < len(dp[0]); j++ { // amount
//             coinVal := coins[i-1]
//             amount := j
//             if coinVal > amount {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = int(math.Min(float64(dp[i][j-coinVal]+1),float64(dp[i-1][j])))
//             }
//         }
//     }
//     result := dp[len(dp)-1][len(dp[0])-1]
//     if result == 10001 {
//         return -1
//     }
//     return result
// }

func coinChange(coins []int, amount int) int {
    m := len(coins)
    n := amount
    
    dp := make([]int, n+1)
    for i := 1; i < len(dp); i++ {
        dp[i] = amount+1 
    }
    
    for i := 1; i < m+1; i++ {
        for j := 1; j < len(dp); j++ {
            coinVal := coins[i-1]
            amount := j
            if coinVal > amount {continue}
            dp[j] = int(math.Min(float64(dp[j]),float64(dp[j-coinVal]+1)))
        }
    }
    
    if dp[len(dp)-1] > amount {return -1}
    return dp[len(dp)-1]
}