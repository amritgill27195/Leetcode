func solveNQueens(n int) [][]string {
    result := [][]string{}
    
    board := make([][]int, n)
    for idx, _ := range board {
        board[idx] = make([]int, n)
    }
    
    var dfs func(r int)
    dfs = func(r int) {
        // base
        if r == n {
            // found a placement combination that is safe for all n queens
            temp := []string{}
            for i := 0; i < n; i++ {
                rowData := new(strings.Builder)
                for j := 0; j < n; j++ {
                    if board[i][j] == 1 {
                        rowData.WriteString("Q")
                    } else {
                        rowData.WriteString(".")
                    }
                }
                temp = append(temp, rowData.String())
            }
            result = append(result, temp)
            return
        }
        
        // logic
        for i := 0; i < n; i++ {
            if isSafe(r, i, board) {
                // action
                board[r][i] = 1
                // recurse
                dfs(r+1)
                // backtrack
                board[r][i] = 0
            }
        }
    }
    dfs(0)
    return result
}

func isSafe(r, c int, board [][]int) bool {
    dirs := [][]int{{-1, 0}, {-1,-1}, {-1,1}}
    n := len(board)
    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        for nr >= 0 && nc >= 0 && nc < n {
            if board[nr][nc] == 1 {
                return false
            }
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}