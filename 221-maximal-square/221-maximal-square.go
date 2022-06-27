func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    max := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                l := 1
                flag := true
                for i + l < m && j + l < n && flag {
                    
                    for k := i+l; k >= i; k--{
                        if matrix[k][j+l] == '0' {
                            flag = false
                            break
                        }
                    }
                    for k := j+l; k >= j; k--{
                        if matrix[i+l][k] == '0' {
                            flag = false
                            break
                        }
                    }
                    
                    if flag { l++ }
                }
                if l*l > max { max = l*l }
            }
        }
    }
    return max
}