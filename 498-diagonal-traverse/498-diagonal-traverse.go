func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    up := true

    r := 0
    c := 0
    idx := 0
    
    result := []int{}
    for idx < m*n {
        result = append(result, mat[r][c])
        idx++
        if up {
            if r == 0 || c == n-1 {
                if c != n-1 {
                    c++
                } else {
                    r++
                }
                up = false
            } else {
                r--
                c++
            }
        } else {
            if r == m-1 || c == 0 {
                if r != m-1 {
                    r++
                } else {
                    c++
                }
                up = true
            } else {
                r++
                c--
            }
        }

    }
    return result
}