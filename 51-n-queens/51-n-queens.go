func solveNQueens(n int) [][]string {
    result := [][]string{}
    
    // generate board
    board := make([][]bool, n)
    for idx, _ := range board{
        board[idx] = make([]bool, n)
    }
    
    var backtrack func(row int)
    backtrack = func(row int) {
        // base 
        if row == n {
            tmp := []string{}
            for i := 0; i < n; i++ {
                rowStr := ""
                for j := 0; j < n; j++ {
                    if board[i][j] {
                        rowStr += "Q"
                    } else {
                        rowStr += "."
                    }
                }
                tmp = append(tmp, rowStr)
            }
            result = append(result, tmp)
            return
        }
        
        
        // logic
        for i := 0; i < n; i++ {
            if itsSafe(row, i, board) {
                board[row][i] = true
                backtrack(row+1)
                board[row][i] = false
            }
        }
    }
    backtrack(0)
    return result
}

func itsSafe(r, c int, board [][]bool) bool {
    dirs := [][]int{
        {-1,0},
        {-1,-1},
        {-1,1},
    }
    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        for nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board) {
            if board[nr][nc]  == true {
                return false
            }
            nr += dir[0]
            nc += dir[1]
        }
    }
    
    return true
}