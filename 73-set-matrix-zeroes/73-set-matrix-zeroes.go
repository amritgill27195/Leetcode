func setZeroes(matrix [][]int)  { 
    m := len(matrix)
    n := len(matrix[0])
    
    cp := make([][]int, m)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if cp[i] == nil {
                cp[i] = make([]int, n)
            } 
            cp[i][j] = math.MaxInt64
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                
                for c := 0; c < n; c++ {
                    cp[i][c] = 0
                }
                
                for c := 0; c < m; c++ {
                    cp[c][j] = 0
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if cp[i][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    
}