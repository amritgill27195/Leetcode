func updateBoard(board [][]byte, click []int) [][]byte {
    if board == nil {
        return nil
    }
    cr := click[0]
    cc := click[1]
    if board[cr][cc] == 'M' {
        board[cr][cc] = 'X'
        return board
    }
    m := len(board)
    n := len(board[0])
    dirs := [][]int{
        {-1,0},{1,0},{0,-1},{0,1},
        {-1,-1},{-1,1},{1,-1},{1,1},
    }
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != 'E' {
            return 
        }
        
        // logic
        numMines := countMines(r,c,board,dirs)
        if numMines > 0 {
            board[r][c] = byte(numMines+'0')
        } else {
            board[r][c] = 'B'
            for _, dir := range dirs {
                dfs(r+dir[0], c+dir[1])
            }
        }
    }
    dfs(cr, cc)
    return board
}


func countMines(r, c int, board [][]byte, dirs [][]int) int {
    count := 0
    m := len(board)
    n := len(board[0])

    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'M' {
            count++
        }
    }
    return count
}