func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    
    r := m-1
    c := 0
    
    for r >= 0 && r < m && c >= 0 && c < n {
        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] > target {
            r--
        } else {
            c++
        }
    }
    return false
}