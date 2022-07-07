func exist(board [][]byte, word string) bool {

    if board == nil || len(board) == 0 || len(word) == 0 { return true }
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    var dfs func(r, c, ptr int) bool
    dfs = func(r, c, ptr int) bool{
        // base
        if ptr == len(word) {
            return true
        }
        if r == m || r < 0 || c == n || c < 0 || board[r][c] == '#' || board[r][c] != word[ptr] {
            return false
        }
        
        // logic
        tmp := board[r][c]
        board[r][c] = '#'
        for _, dir := range dirs {
            if dfs(r+dir[0], c+dir[1], ptr+1) {
                return true
            }
        }
        board[r][c] = tmp
        
        return false
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                if found := dfs(i, j, 0); found {
                    return true
                }
            }
        }
    }
    return false
}