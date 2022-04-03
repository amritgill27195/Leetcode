func isValidSudoku(board [][]byte) bool {
    
    m := len(board)
    n := len(board[0])
    boxSet := make([]map[byte]struct{}, m)
    for i := 0; i < m; i++ {
        rowSet := map[byte]struct{}{}
        for j := 0; j < n; j++ {
            if board[i][j] == '.' {continue}
            val := board[i][j]
            
            if _, ok := rowSet[val]; ok {
                return false
            }
            boxIdx := (i/3) * 3 + (j/3)
            if _, ok := boxSet[boxIdx][val]; ok {
                return false
            }
            rowSet[board[i][j]] = struct{}{}
            if boxSet[boxIdx] == nil {boxSet[boxIdx] = map[byte]struct{}{}}
            boxSet[boxIdx][val] = struct{}{}
        }
    }
    
   
    for i := 0; i < m; i++ {
        colSet := map[byte]struct{}{}
        for j := 0; j < n; j++ {
            if string(board[j][i]) == "." {continue}
            if _, ok := colSet[board[j][i]]; ok {
                return false
            }
            colSet[board[j][i]] = struct{}{}
        }
    }
    
    return true
    
}