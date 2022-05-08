func generateMatrix(n int) [][]int {
 
    if n == 0 {
        return nil
    }
    
    
    total := n * n
    mx := make([][]int, n)
    for idx, _ := range mx {
        mx[idx] = make([]int, n)
    }
    
    left := 0
    top := 0
    right := n-1
    bottom := n-1
    
    c := 1
    for c <= total {
        
        // top row
        // left to right
        for x := left; x <= right; x++ {
            mx[top][x] = c
            c++
        }
        top++
        
        
        // right col
        // top to bottom
        for x := top; x <= bottom; x++ {
            mx[x][right] = c
            c++
        }
        right--
        
        // bottom row
        // right to left
        for x := right; x >= left; x-- {
            mx[bottom][x] = c
            c++
        }
        bottom--
        
        // left col
        // bottom to top
        for x := bottom; x >= top; x-- {
            mx[x][left] = c
            c++
        }
        left++
    }
    return mx
}