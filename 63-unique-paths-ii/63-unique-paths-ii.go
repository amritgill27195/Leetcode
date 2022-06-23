func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    if obstacleGrid[m-1][n-1] == 1 ||  obstacleGrid[0][0] == 1 {
        return 0
    }
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    dp[m-1][n-1] = 1
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i == m-1 && j == n-1 {continue}
            br := i+1
            bc := j
            rr := i
            rc := j+1
            bottomVal := dp[i+1][j]
            if br < m && br >= 0 && bc < n && bc >= 0 && obstacleGrid[br][bc] == 1 {
                bottomVal = 0
            }
            
            rightVal := dp[i][j+1]
            if rr < m && rr >= 0 && rc < n && rc >= 0 && obstacleGrid[rr][rc] == 1 {
                rightVal = 0
            }
            
            dp[i][j] = bottomVal + rightVal
            
        }
    }
    return dp[0][0]
}