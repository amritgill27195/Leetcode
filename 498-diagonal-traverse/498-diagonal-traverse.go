func findDiagonalOrder(mat [][]int) []int {
    if mat == nil || len(mat) == 0 {
        return nil
    }
    m := len(mat)
    n := len(mat[0])
    out := []int{}
    idx := 0
    r := 0
    c := 0
    up := true
    for idx < m*n {
        out = append(out, mat[r][c])
        idx++
        
        if up {
            if c == n-1 {
                up = false
                r++
            } else if r == 0 {
                up = false
                c++
            }  else {
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
            }  else {
                r++
                c--
            }
        
        }
        
        
    }
    
    
    return out
}