func diagonalSum(mat [][]int) int {
    r := 0
    c := 0
    m := len(mat)
    n := len(mat[0])
    sum := 0
    
    for r < m && c < n {
        sum += mat[r][c]
        mat[r][c] *= -1
        r++
        c++
    }
    
    r = 0
    c = n-1
    
    for r < m && c >= 0 {
        if mat[r][c] > 0 {
            sum += mat[r][c]
        }
        r++
        c--
    }
    
    return sum
}