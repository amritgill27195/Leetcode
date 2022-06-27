// func maximalSquare(matrix [][]byte) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     max := 0
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == '1' {
//                 l := 1
//                 flag := true
//                 for i + l < m && j + l < n && flag {
                    
//                     for k := i+l; k >= i; k--{
//                         if matrix[k][j+l] == '0' {
//                             flag = false
//                             break
//                         }
//                     }
//                     for k := j+l; k >= j; k--{
//                         if matrix[i+l][k] == '0' {
//                             flag = false
//                             break
//                         }
//                     }
                    
//                     if flag { l++ }
//                 }
//                 max = int(math.Max(float64(max),float64(l*l)))
//             }
//         }
//     }
//     return max
// }


func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int,n+1)
    }
    max := 0
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if matrix[i-1][j-1] == '1' {
                min := int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))))
                dp[i][j] = min+1
                if dp[i][j] > max {
                    max = dp[i][j]
                }
            }
        }
    }
    
    return max*max
}