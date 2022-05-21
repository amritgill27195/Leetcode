func exist(board [][]byte, word string) bool {
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m := len(board)
    n := len(board[0])
    var backtrack func(r, c int, wordIdx int) bool
    backtrack = func(r, c int, wordIdx int) bool {
        // base
        if wordIdx == len(word) {return true}
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == '#' || board[r][c] != word[wordIdx] {
            return false
        } 
        
        // logic
        tmp := board[r][c]
        board[r][c] = '#'
        for _, dir := range dirs {
            if found := backtrack(r+dir[0],c+dir[1], wordIdx+1); found {
                return true
            }
        }
        board[r][c] = tmp
        return false
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                if found := backtrack(i,j,0); found {
                    return true
                }
            }
        }
    }
    
    return false
    
}