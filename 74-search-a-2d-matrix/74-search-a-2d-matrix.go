func searchMatrix(matrix [][]int, target int) bool {
    
    if matrix == nil || len(matrix) == 0 {
        return false
    }
    
    outterSize := len(matrix)
    innerSize := len(matrix[0])
    totalLenOfImaginary1D := outterSize * innerSize
    
    left := 0
    right := totalLenOfImaginary1D
    for left <= right {
        mid := left + ((right-left)/2)
        outterIdx := mid/innerSize
        innerIdx := mid%innerSize
    
        // I do not like this way of handling.... 
        // TODO: find a better clean way to handle- IndexOutOfBounds
        if outterIdx >= outterSize || innerIdx >= innerSize {
            break
        }
        
        midVal := matrix[outterIdx][innerIdx]
        if target == midVal {
            return true
        } else if target < midVal {
            right = mid-1
        } else {
            left = mid+1
        }        
    }
    return false
}