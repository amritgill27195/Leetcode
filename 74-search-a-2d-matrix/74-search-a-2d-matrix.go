func searchMatrix(matrix [][]int, target int) bool {
    numRows := len(matrix)
    numCols := len(matrix[0])
    total := numRows * numCols
    
    left := 0
    right := total-1
    for left <= right {
        mid := left + (right-left)/2
        row := mid / numCols
        col := mid % numCols
        if row >= numRows || col >= numCols {
            break
        }
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    
    return false
}