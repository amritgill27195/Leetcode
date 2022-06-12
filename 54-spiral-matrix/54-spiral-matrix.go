func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    
    left := 0
    right := n-1
    top := 0
    bottom := m-1
    result := []int{}
    for left <= right && top <= bottom {
        
        // left to right
        for i := top; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++
        
        
        if left <= right && top <= bottom{ 
            for i := top; i <= bottom; i++ {
                result = append(result, matrix[i][right])
            }
        }
        right--

        
        if left <= right && top <= bottom {
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
        }
        bottom--
        
        
        if left <= right && top <= bottom{
            for i := bottom; i>=top; i-- {
                result = append(result, matrix[i][left])
            }
        }
        left++
    
    }
    return result
}
