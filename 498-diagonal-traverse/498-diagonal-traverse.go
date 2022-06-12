func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    result := []int{}
    r := 0
    c := 0
    up := true
    idx := 0
    
    for idx < m*n {
        result = append(result, mat[r][c])
        idx++
        if up {
            if c == n-1 {
                r++
                up = false
            } else if r == 0 {
                c++
                up = false
            } else {
                r--
                c++
            }
        } else {
            if r == m-1 {
                c++
                up = true
            } else if c == 0 {
                r++
                up = true
            } else {
                r++
                c--
            }
        }
    }
    return result
}