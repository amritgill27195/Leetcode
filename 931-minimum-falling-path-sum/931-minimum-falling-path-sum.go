// time: o(mn)
// space: o(1)
// func minFallingPathSum(matrix [][]int) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     for i := m-2; i >= 0; i-- {
//         for j := 0; j < n; j++ {
//             if j == 0 {
//                 matrix[i][j] += int(math.Min(float64(matrix[i+1][0]), float64(matrix[i+1][1])))
//             } else if j == n-1 {
//                 matrix[i][j] += int(math.Min(float64(matrix[i+1][j]), float64(matrix[i+1][j-1])))
//             } else {
//                 matrix[i][j] += int(math.Min(float64(matrix[i+1][j]), math.Min(float64(matrix[i+1][j-1]),float64(matrix[i+1][j+1]))))
//             }
//         }
//     }
//     min := matrix[0][0]
//     for j := 1; j < n; j++ {
//         if matrix[0][j] < min {
//             min = matrix[0][j]
//         }
//     }
//     return min
// }


func minFallingPathSum(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    for idx, _ := range dp {
        dp[idx] = make([]int, n)
    }
    for j := 0; j < n; j++ {
        dp[m-1][j] = matrix[m-1][j]
    }
    
    for i := m-2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            if j == 0 {
                dp[i][j] = matrix[i][j] + int(math.Min(float64(dp[i+1][0]), float64(dp[i+1][1])))
            } else if j == n-1 {
                dp[i][j] = matrix[i][j] + int(math.Min(float64(dp[i+1][j]), float64(dp[i+1][j-1])))
            } else {
                dp[i][j] = matrix[i][j] + int(math.Min(float64(dp[i+1][j]), math.Min(float64(dp[i+1][j-1]),float64(dp[i+1][j+1]))))
            }
        }
    }

    min := dp[0][0]
    for j := 1; j < n; j++ {
        if dp[0][j] < min {
            min = dp[0][j]
        }
    }
    return min
}