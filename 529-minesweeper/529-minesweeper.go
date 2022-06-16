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
    board[cr][cc] = 'B'
    q := [][]int{{cr,cc}}
    dirs := [][]int{
        {-1,0},
        {1,0},
        {0,-1},
        {0,1},
        {-1,-1},
        {-1,1},
        {1,-1},
        {1,1},
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        currRow := dq[0]
        currCol := dq[1]
        numMines := countMines(currRow, currCol, board, dirs)
        if numMines > 0 {
            board[currRow][currCol] = byte(numMines + '0')
        } else {
            for _, dir := range dirs {
                nr := currRow+dir[0]
                nc := currCol+dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'E' {
                    board[nr][nc] = 'B'
                    q = append(q, []int{nr, nc})
                }
            }
        }
    }
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