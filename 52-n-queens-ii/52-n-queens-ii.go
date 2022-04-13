func totalNQueens(n int) int {
      // generate the nxn board
    board := make([][]bool, n)
    for idx, _ := range board {
        board[idx] = make([]bool, n)
    }
    
    count := 0
    // now try placing queens 
    var backtrack func(row int)
    backtrack = func(row int) {
        // base
        if row == n {
            count++
            return
        }
        
        //logic
        // we loop over cols, row is controlled by recursion
        for i := 0; i < n; i++ {
            if canBePlaced(row, i, board) {
                // action
                board[row][i] = true
                // recurse
                backtrack(row+1)
                // backtrack
                board[row][i] = false
            }
        }
        
    }
    backtrack(0)
    return count
}


// since we are looping row by row
// the only directions we have to check for a specific [r][c]
// are vertical up, diagonal up left and diagonal up right
// we do not have to check in the same row because this function will be called when we do not have anything placed yet
// we also do not have to check anything below because we have not went to row+1 yet,
// and so we only need to check up from a specific r,c
func canBePlaced(r, c int, board [][]bool) bool {
    n := len(board)
    dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
    for _, dir := range dirs {
        newR := r + dir[0]
        newC := c + dir[1]
        for newR >= 0 && newR < n && newC >= 0 && newC < n {
            if board[newR][newC] == true {
                return false
            }
            // continue in the same direction as long as we are inbound
            newR = newR + dir[0]
            newC = newC + dir[1]
        }  
    }
    return true
} 
