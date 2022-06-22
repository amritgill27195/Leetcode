func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    top := 0
    right := n-1
    bottom := m-1
    left := 0
    
    out := make([]int, m*n)
    idx := 0
    for top <= bottom && left <= right {
        // left to right
        for i := left; i <= right; i++ {
            out[idx] = matrix[left][i]
            idx++
        }
        top++
        
        if top <= bottom && left <= right {
            // top to bottom
            for i := top; i <= bottom; i++ {
                out[idx] = matrix[i][right]
                idx++
            }
        }
        right--
        
        if top <= bottom && left <= right {
            // bottom row right to left
            for i := right; i >= left; i-- {
                out[idx] = matrix[bottom][i]
                idx++
            }
        }
        bottom--
        
        if top <= bottom && left <= right {
            // left col - bottom to top
            for i := bottom; i >= top; i-- {
                out[idx] = matrix[i][left]
                idx++
            }
        }
        left++
    }
    return out
    
}