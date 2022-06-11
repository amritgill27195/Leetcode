func minFallingPathSum(matrix [][]int) int {
    dirs := [][]int{{1,0}, {1,-1}, {1,1}}
    m := len(matrix)
    n := len(matrix[0])
    
    for i := m-2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            min := math.MaxInt64
            for _, dir := range dirs {
                nr := i+dir[0]
                nc := j+dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n {
                    if matrix[nr][nc] < min {
                        min = matrix[nr][nc]
                    }
                }
            }
            matrix[i][j] += min
        }
    }
    minVal := matrix[0][0]
    for j := 1; j < n; j++ {
        if matrix[0][j] < minVal {
            minVal = matrix[0][j]
        }
    }
    return minVal    
}