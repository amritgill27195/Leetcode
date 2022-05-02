// brute force recursive solution
// time: o(2^amount)
// func coinChange(coins []int, amount int) int {
//     var helper func(startIdx int, target int, numCoinsUsed int) int
//     helper = func(startIdx int, target int, numCoinsUsed int) int {
//         // base
//         if target == 0 {
//             return numCoinsUsed
//         }
//         if startIdx == len(coins) || target < 0 {
//             return -1
//         }
        
//         // logic
//         choose := helper(startIdx, target-coins[startIdx], numCoinsUsed+1)
//         notChoose := helper(startIdx+1, target, numCoinsUsed)
//         if choose == -1 || notChoose == -1 {
//             return int(math.Max(float64(choose), float64(notChoose)))
//         }
        
//         return int(math.Min(float64(choose),float64(notChoose)))
//     }
//     return helper(0, amount, 0)
// }


func coinChange(coins []int, amount int) int {
    m := len(coins)
    n := amount
    
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    for c := 1; c < n+1; c++ {
        dp[0][c] = amount+1
    }
        
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            coinVal := coins[i-1]
            am := j
            
            if coinVal > am {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = int(math.Min(float64(dp[i][j-coinVal]+1),float64(dp[i-1][j])))
            }
        }
    }
    result := dp[len(dp)-1][len(dp[0])-1]
    if result > amount {
        return -1
    }
    return result
}