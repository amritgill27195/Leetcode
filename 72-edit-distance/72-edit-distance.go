func minDistance(word1 string, word2 string) int {
    
    if word1 == word2 {
        return 0
    }
    m := len(word2)
    n := len(word1)
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    for j := 0; j < len(dp[0]); j++ {  dp[0][j] = j  }
    for i := 0; i < len(dp); i++    { dp[i][0] = i }

    for i := 1; i < len(dp); i++ {  // word2
        for j := 1; j < len(dp[0]); j++ { // word1
            word1Char := word1[j-1]
            word2Char := word2[i-1]
            if word1Char == word2Char {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1          
            }
        }
    }
    
    return dp[len(dp)-1][len(dp[0])-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}