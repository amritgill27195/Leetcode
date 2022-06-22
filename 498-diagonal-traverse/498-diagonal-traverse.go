func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    out := make([]int, m*n)
    idx := 0
    r := 0
    c := 0
    up := true
    for idx < m*n {
        out[idx] = mat[r][c]
        idx++
        if up {
            if r == 0 || c == n-1 {
                if c == n-1 {
                    r++
                } else {
                    c++
                }
                up = false
            } else {
                r--;c++
            }
        } else {
            if c == 0 || r == m-1 {
                if r == m-1 {
                    c++
                } else {
                    r++
                }
                up = true
            } else {
                r++; c--
            }
        }
    }
    return out
}