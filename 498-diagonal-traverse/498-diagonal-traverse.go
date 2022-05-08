func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    
    r := 0
    c := 0
    idx := 0
    up := true
    
    var out []int
    for idx < m*n {
        idx++
        out = append(out, mat[r][c])
        
        if up {
            if c == n-1 {
                up = false
                r++
            } else if r == 0 {
                up = false
                c++
            } else {
                r--
                c++
            }
        } else {
            if r == m-1 {
                up = true
                c++
            } else if c == 0 {
                up = true
                r++
            } else {
                r++
                c--
            }
        }
    }
    return out
}